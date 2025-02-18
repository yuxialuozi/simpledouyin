package feed

import (
	"context"
	"simpledouyin/config"
	bizConstant "simpledouyin/constants/biz"
	"simpledouyin/kitex_gen/douyin/feed"
	feedService "simpledouyin/kitex_gen/douyin/feed/feedservice"
	"simpledouyin/logging"
	"simpledouyin/service/web/mw"
	"strconv"

	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"

	"github.com/cloudwego/hertz/pkg/app"
	httpStatus "github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/sirupsen/logrus"
)

var feedClient feedService.Client

func init() {
	r, err := consul.NewConsulResolver(config.EnvConfig.CONSUL_ADDR)
	if err != nil {
		logging.Logger.WithError(err).Fatal("init feed client failed")
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.FeedServiceName),
		provider.WithExportEndpoint(config.EnvConfig.EXPORT_ENDPOINT),
		provider.WithInsecure(),
	)
	feedClient, err = feedService.NewClient(
		config.FeedServiceName,
		client.WithResolver(r),
		client.WithSuite(tracing.NewClientSuite()),
	)
	if err != nil {
		logging.Logger.WithError(err).Fatal("init feed client failed")
		panic(err)
	}
}
func Action(ctx context.Context, c *app.RequestContext) {
	methodFields := logrus.Fields{
		"method": "FeedAction",
	}
	logger := logging.Logger.WithFields(methodFields)
	logger.Debugf("Process start")

	latestTime := c.Query("latest_time")
	if _, err := strconv.Atoi(latestTime); latestTime != "" && err != nil {
		bizConstant.InvalidLatestTime.
			WithCause(err).
			WithFields(&methodFields).
			LaunchError(c)
		return
	}

	actorIdPtr, ok := mw.Auth(c)
	actorId := *actorIdPtr
	if !ok {
		return
	}

	logger.WithFields(logrus.Fields{
		"latestTime": latestTime,
		"actorId":    actorId,
	}).Debugf("Executing get feed")
	response, err := feedClient.ListVideos(ctx, &feed.ListFeedRequest{
		LatestTime: &latestTime,
		ActorId:    actorIdPtr,
	})

	if err != nil {
		bizConstant.RPCCallError.
			WithCause(err).
			WithFields(&methodFields).
			LaunchError(c)
		return
	}

	logger.WithFields(logrus.Fields{
		"response": response,
	}).Debugf("Getting feed success")
	c.JSON(httpStatus.StatusOK, response)
}
