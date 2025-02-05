// Code generated by Kitex v0.12.1. DO NOT EDIT.

package feedservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	feed "simpledouyin/kitex_gen/douyin/feed"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"ListVideos": kitex.NewMethodInfo(
		listVideosHandler,
		newListVideosArgs,
		newListVideosResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"QueryVideos": kitex.NewMethodInfo(
		queryVideosHandler,
		newQueryVideosArgs,
		newQueryVideosResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	feedServiceServiceInfo                = NewServiceInfo()
	feedServiceServiceInfoForClient       = NewServiceInfoForClient()
	feedServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return feedServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return feedServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return feedServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "FeedService"
	handlerType := (*feed.FeedService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "simpledouyin.feed",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.12.1",
		Extra:           extra,
	}
	return svcInfo
}

func listVideosHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(feed.ListFeedRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(feed.FeedService).ListVideos(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ListVideosArgs:
		success, err := handler.(feed.FeedService).ListVideos(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListVideosResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newListVideosArgs() interface{} {
	return &ListVideosArgs{}
}

func newListVideosResult() interface{} {
	return &ListVideosResult{}
}

type ListVideosArgs struct {
	Req *feed.ListFeedRequest
}

func (p *ListVideosArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(feed.ListFeedRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListVideosArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListVideosArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListVideosArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ListVideosArgs) Unmarshal(in []byte) error {
	msg := new(feed.ListFeedRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListVideosArgs_Req_DEFAULT *feed.ListFeedRequest

func (p *ListVideosArgs) GetReq() *feed.ListFeedRequest {
	if !p.IsSetReq() {
		return ListVideosArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListVideosArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ListVideosArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ListVideosResult struct {
	Success *feed.ListFeedResponse
}

var ListVideosResult_Success_DEFAULT *feed.ListFeedResponse

func (p *ListVideosResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(feed.ListFeedResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListVideosResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListVideosResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListVideosResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ListVideosResult) Unmarshal(in []byte) error {
	msg := new(feed.ListFeedResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListVideosResult) GetSuccess() *feed.ListFeedResponse {
	if !p.IsSetSuccess() {
		return ListVideosResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListVideosResult) SetSuccess(x interface{}) {
	p.Success = x.(*feed.ListFeedResponse)
}

func (p *ListVideosResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListVideosResult) GetResult() interface{} {
	return p.Success
}

func queryVideosHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(feed.QueryVideosRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(feed.FeedService).QueryVideos(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *QueryVideosArgs:
		success, err := handler.(feed.FeedService).QueryVideos(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QueryVideosResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newQueryVideosArgs() interface{} {
	return &QueryVideosArgs{}
}

func newQueryVideosResult() interface{} {
	return &QueryVideosResult{}
}

type QueryVideosArgs struct {
	Req *feed.QueryVideosRequest
}

func (p *QueryVideosArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(feed.QueryVideosRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *QueryVideosArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *QueryVideosArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *QueryVideosArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *QueryVideosArgs) Unmarshal(in []byte) error {
	msg := new(feed.QueryVideosRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryVideosArgs_Req_DEFAULT *feed.QueryVideosRequest

func (p *QueryVideosArgs) GetReq() *feed.QueryVideosRequest {
	if !p.IsSetReq() {
		return QueryVideosArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryVideosArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *QueryVideosArgs) GetFirstArgument() interface{} {
	return p.Req
}

type QueryVideosResult struct {
	Success *feed.QueryVideosResponse
}

var QueryVideosResult_Success_DEFAULT *feed.QueryVideosResponse

func (p *QueryVideosResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(feed.QueryVideosResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *QueryVideosResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *QueryVideosResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *QueryVideosResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *QueryVideosResult) Unmarshal(in []byte) error {
	msg := new(feed.QueryVideosResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryVideosResult) GetSuccess() *feed.QueryVideosResponse {
	if !p.IsSetSuccess() {
		return QueryVideosResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryVideosResult) SetSuccess(x interface{}) {
	p.Success = x.(*feed.QueryVideosResponse)
}

func (p *QueryVideosResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *QueryVideosResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ListVideos(ctx context.Context, Req *feed.ListFeedRequest) (r *feed.ListFeedResponse, err error) {
	var _args ListVideosArgs
	_args.Req = Req
	var _result ListVideosResult
	if err = p.c.Call(ctx, "ListVideos", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryVideos(ctx context.Context, Req *feed.QueryVideosRequest) (r *feed.QueryVideosResponse, err error) {
	var _args QueryVideosArgs
	_args.Req = Req
	var _result QueryVideosResult
	if err = p.c.Call(ctx, "QueryVideos", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
