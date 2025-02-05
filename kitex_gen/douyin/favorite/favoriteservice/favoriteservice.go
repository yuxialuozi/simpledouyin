// Code generated by Kitex v0.12.1. DO NOT EDIT.

package favoriteservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	favorite "simpledouyin/kitex_gen/douyin/favorite"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"FavoriteAction": kitex.NewMethodInfo(
		favoriteActionHandler,
		newFavoriteActionArgs,
		newFavoriteActionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"FavoriteList": kitex.NewMethodInfo(
		favoriteListHandler,
		newFavoriteListArgs,
		newFavoriteListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"IsFavorite": kitex.NewMethodInfo(
		isFavoriteHandler,
		newIsFavoriteArgs,
		newIsFavoriteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"CountFavorite": kitex.NewMethodInfo(
		countFavoriteHandler,
		newCountFavoriteArgs,
		newCountFavoriteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"CountUserFavorite": kitex.NewMethodInfo(
		countUserFavoriteHandler,
		newCountUserFavoriteArgs,
		newCountUserFavoriteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"CountUserTotalFavorited": kitex.NewMethodInfo(
		countUserTotalFavoritedHandler,
		newCountUserTotalFavoritedArgs,
		newCountUserTotalFavoritedResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	favoriteServiceServiceInfo                = NewServiceInfo()
	favoriteServiceServiceInfoForClient       = NewServiceInfoForClient()
	favoriteServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return favoriteServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return favoriteServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return favoriteServiceServiceInfoForClient
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
	serviceName := "FavoriteService"
	handlerType := (*favorite.FavoriteService)(nil)
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
		"PackageName": "simpledouyin.favorite",
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

func favoriteActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favorite.FavoriteRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favorite.FavoriteService).FavoriteAction(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *FavoriteActionArgs:
		success, err := handler.(favorite.FavoriteService).FavoriteAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FavoriteActionResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newFavoriteActionArgs() interface{} {
	return &FavoriteActionArgs{}
}

func newFavoriteActionResult() interface{} {
	return &FavoriteActionResult{}
}

type FavoriteActionArgs struct {
	Req *favorite.FavoriteRequest
}

func (p *FavoriteActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(favorite.FavoriteRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FavoriteActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FavoriteActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FavoriteActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *FavoriteActionArgs) Unmarshal(in []byte) error {
	msg := new(favorite.FavoriteRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FavoriteActionArgs_Req_DEFAULT *favorite.FavoriteRequest

func (p *FavoriteActionArgs) GetReq() *favorite.FavoriteRequest {
	if !p.IsSetReq() {
		return FavoriteActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FavoriteActionArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *FavoriteActionArgs) GetFirstArgument() interface{} {
	return p.Req
}

type FavoriteActionResult struct {
	Success *favorite.FavoriteResponse
}

var FavoriteActionResult_Success_DEFAULT *favorite.FavoriteResponse

func (p *FavoriteActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(favorite.FavoriteResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FavoriteActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FavoriteActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FavoriteActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *FavoriteActionResult) Unmarshal(in []byte) error {
	msg := new(favorite.FavoriteResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FavoriteActionResult) GetSuccess() *favorite.FavoriteResponse {
	if !p.IsSetSuccess() {
		return FavoriteActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FavoriteActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*favorite.FavoriteResponse)
}

func (p *FavoriteActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FavoriteActionResult) GetResult() interface{} {
	return p.Success
}

func favoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favorite.FavoriteListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favorite.FavoriteService).FavoriteList(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *FavoriteListArgs:
		success, err := handler.(favorite.FavoriteService).FavoriteList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FavoriteListResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newFavoriteListArgs() interface{} {
	return &FavoriteListArgs{}
}

func newFavoriteListResult() interface{} {
	return &FavoriteListResult{}
}

type FavoriteListArgs struct {
	Req *favorite.FavoriteListRequest
}

func (p *FavoriteListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(favorite.FavoriteListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FavoriteListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FavoriteListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FavoriteListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *FavoriteListArgs) Unmarshal(in []byte) error {
	msg := new(favorite.FavoriteListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FavoriteListArgs_Req_DEFAULT *favorite.FavoriteListRequest

func (p *FavoriteListArgs) GetReq() *favorite.FavoriteListRequest {
	if !p.IsSetReq() {
		return FavoriteListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FavoriteListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *FavoriteListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type FavoriteListResult struct {
	Success *favorite.FavoriteListResponse
}

var FavoriteListResult_Success_DEFAULT *favorite.FavoriteListResponse

func (p *FavoriteListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(favorite.FavoriteListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FavoriteListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FavoriteListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FavoriteListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *FavoriteListResult) Unmarshal(in []byte) error {
	msg := new(favorite.FavoriteListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FavoriteListResult) GetSuccess() *favorite.FavoriteListResponse {
	if !p.IsSetSuccess() {
		return FavoriteListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FavoriteListResult) SetSuccess(x interface{}) {
	p.Success = x.(*favorite.FavoriteListResponse)
}

func (p *FavoriteListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FavoriteListResult) GetResult() interface{} {
	return p.Success
}

func isFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favorite.IsFavoriteRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favorite.FavoriteService).IsFavorite(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *IsFavoriteArgs:
		success, err := handler.(favorite.FavoriteService).IsFavorite(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*IsFavoriteResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newIsFavoriteArgs() interface{} {
	return &IsFavoriteArgs{}
}

func newIsFavoriteResult() interface{} {
	return &IsFavoriteResult{}
}

type IsFavoriteArgs struct {
	Req *favorite.IsFavoriteRequest
}

func (p *IsFavoriteArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(favorite.IsFavoriteRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *IsFavoriteArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *IsFavoriteArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *IsFavoriteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *IsFavoriteArgs) Unmarshal(in []byte) error {
	msg := new(favorite.IsFavoriteRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var IsFavoriteArgs_Req_DEFAULT *favorite.IsFavoriteRequest

func (p *IsFavoriteArgs) GetReq() *favorite.IsFavoriteRequest {
	if !p.IsSetReq() {
		return IsFavoriteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *IsFavoriteArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *IsFavoriteArgs) GetFirstArgument() interface{} {
	return p.Req
}

type IsFavoriteResult struct {
	Success *favorite.IsFavoriteResponse
}

var IsFavoriteResult_Success_DEFAULT *favorite.IsFavoriteResponse

func (p *IsFavoriteResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(favorite.IsFavoriteResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *IsFavoriteResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *IsFavoriteResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *IsFavoriteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *IsFavoriteResult) Unmarshal(in []byte) error {
	msg := new(favorite.IsFavoriteResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *IsFavoriteResult) GetSuccess() *favorite.IsFavoriteResponse {
	if !p.IsSetSuccess() {
		return IsFavoriteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *IsFavoriteResult) SetSuccess(x interface{}) {
	p.Success = x.(*favorite.IsFavoriteResponse)
}

func (p *IsFavoriteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *IsFavoriteResult) GetResult() interface{} {
	return p.Success
}

func countFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favorite.CountFavoriteRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favorite.FavoriteService).CountFavorite(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CountFavoriteArgs:
		success, err := handler.(favorite.FavoriteService).CountFavorite(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CountFavoriteResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCountFavoriteArgs() interface{} {
	return &CountFavoriteArgs{}
}

func newCountFavoriteResult() interface{} {
	return &CountFavoriteResult{}
}

type CountFavoriteArgs struct {
	Req *favorite.CountFavoriteRequest
}

func (p *CountFavoriteArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(favorite.CountFavoriteRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CountFavoriteArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CountFavoriteArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CountFavoriteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CountFavoriteArgs) Unmarshal(in []byte) error {
	msg := new(favorite.CountFavoriteRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CountFavoriteArgs_Req_DEFAULT *favorite.CountFavoriteRequest

func (p *CountFavoriteArgs) GetReq() *favorite.CountFavoriteRequest {
	if !p.IsSetReq() {
		return CountFavoriteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CountFavoriteArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CountFavoriteArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CountFavoriteResult struct {
	Success *favorite.CountFavoriteResponse
}

var CountFavoriteResult_Success_DEFAULT *favorite.CountFavoriteResponse

func (p *CountFavoriteResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(favorite.CountFavoriteResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CountFavoriteResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CountFavoriteResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CountFavoriteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CountFavoriteResult) Unmarshal(in []byte) error {
	msg := new(favorite.CountFavoriteResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CountFavoriteResult) GetSuccess() *favorite.CountFavoriteResponse {
	if !p.IsSetSuccess() {
		return CountFavoriteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CountFavoriteResult) SetSuccess(x interface{}) {
	p.Success = x.(*favorite.CountFavoriteResponse)
}

func (p *CountFavoriteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CountFavoriteResult) GetResult() interface{} {
	return p.Success
}

func countUserFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favorite.CountUserFavoriteRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favorite.FavoriteService).CountUserFavorite(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CountUserFavoriteArgs:
		success, err := handler.(favorite.FavoriteService).CountUserFavorite(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CountUserFavoriteResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCountUserFavoriteArgs() interface{} {
	return &CountUserFavoriteArgs{}
}

func newCountUserFavoriteResult() interface{} {
	return &CountUserFavoriteResult{}
}

type CountUserFavoriteArgs struct {
	Req *favorite.CountUserFavoriteRequest
}

func (p *CountUserFavoriteArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(favorite.CountUserFavoriteRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CountUserFavoriteArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CountUserFavoriteArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CountUserFavoriteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CountUserFavoriteArgs) Unmarshal(in []byte) error {
	msg := new(favorite.CountUserFavoriteRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CountUserFavoriteArgs_Req_DEFAULT *favorite.CountUserFavoriteRequest

func (p *CountUserFavoriteArgs) GetReq() *favorite.CountUserFavoriteRequest {
	if !p.IsSetReq() {
		return CountUserFavoriteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CountUserFavoriteArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CountUserFavoriteArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CountUserFavoriteResult struct {
	Success *favorite.CountUserFavoriteResponse
}

var CountUserFavoriteResult_Success_DEFAULT *favorite.CountUserFavoriteResponse

func (p *CountUserFavoriteResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(favorite.CountUserFavoriteResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CountUserFavoriteResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CountUserFavoriteResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CountUserFavoriteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CountUserFavoriteResult) Unmarshal(in []byte) error {
	msg := new(favorite.CountUserFavoriteResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CountUserFavoriteResult) GetSuccess() *favorite.CountUserFavoriteResponse {
	if !p.IsSetSuccess() {
		return CountUserFavoriteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CountUserFavoriteResult) SetSuccess(x interface{}) {
	p.Success = x.(*favorite.CountUserFavoriteResponse)
}

func (p *CountUserFavoriteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CountUserFavoriteResult) GetResult() interface{} {
	return p.Success
}

func countUserTotalFavoritedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favorite.CountUserTotalFavoritedRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favorite.FavoriteService).CountUserTotalFavorited(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CountUserTotalFavoritedArgs:
		success, err := handler.(favorite.FavoriteService).CountUserTotalFavorited(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CountUserTotalFavoritedResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCountUserTotalFavoritedArgs() interface{} {
	return &CountUserTotalFavoritedArgs{}
}

func newCountUserTotalFavoritedResult() interface{} {
	return &CountUserTotalFavoritedResult{}
}

type CountUserTotalFavoritedArgs struct {
	Req *favorite.CountUserTotalFavoritedRequest
}

func (p *CountUserTotalFavoritedArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(favorite.CountUserTotalFavoritedRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CountUserTotalFavoritedArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CountUserTotalFavoritedArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CountUserTotalFavoritedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CountUserTotalFavoritedArgs) Unmarshal(in []byte) error {
	msg := new(favorite.CountUserTotalFavoritedRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CountUserTotalFavoritedArgs_Req_DEFAULT *favorite.CountUserTotalFavoritedRequest

func (p *CountUserTotalFavoritedArgs) GetReq() *favorite.CountUserTotalFavoritedRequest {
	if !p.IsSetReq() {
		return CountUserTotalFavoritedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CountUserTotalFavoritedArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CountUserTotalFavoritedArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CountUserTotalFavoritedResult struct {
	Success *favorite.CountUserTotalFavoritedResponse
}

var CountUserTotalFavoritedResult_Success_DEFAULT *favorite.CountUserTotalFavoritedResponse

func (p *CountUserTotalFavoritedResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(favorite.CountUserTotalFavoritedResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CountUserTotalFavoritedResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CountUserTotalFavoritedResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CountUserTotalFavoritedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CountUserTotalFavoritedResult) Unmarshal(in []byte) error {
	msg := new(favorite.CountUserTotalFavoritedResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CountUserTotalFavoritedResult) GetSuccess() *favorite.CountUserTotalFavoritedResponse {
	if !p.IsSetSuccess() {
		return CountUserTotalFavoritedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CountUserTotalFavoritedResult) SetSuccess(x interface{}) {
	p.Success = x.(*favorite.CountUserTotalFavoritedResponse)
}

func (p *CountUserTotalFavoritedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CountUserTotalFavoritedResult) GetResult() interface{} {
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

func (p *kClient) FavoriteAction(ctx context.Context, Req *favorite.FavoriteRequest) (r *favorite.FavoriteResponse, err error) {
	var _args FavoriteActionArgs
	_args.Req = Req
	var _result FavoriteActionResult
	if err = p.c.Call(ctx, "FavoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteList(ctx context.Context, Req *favorite.FavoriteListRequest) (r *favorite.FavoriteListResponse, err error) {
	var _args FavoriteListArgs
	_args.Req = Req
	var _result FavoriteListResult
	if err = p.c.Call(ctx, "FavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) IsFavorite(ctx context.Context, Req *favorite.IsFavoriteRequest) (r *favorite.IsFavoriteResponse, err error) {
	var _args IsFavoriteArgs
	_args.Req = Req
	var _result IsFavoriteResult
	if err = p.c.Call(ctx, "IsFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CountFavorite(ctx context.Context, Req *favorite.CountFavoriteRequest) (r *favorite.CountFavoriteResponse, err error) {
	var _args CountFavoriteArgs
	_args.Req = Req
	var _result CountFavoriteResult
	if err = p.c.Call(ctx, "CountFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CountUserFavorite(ctx context.Context, Req *favorite.CountUserFavoriteRequest) (r *favorite.CountUserFavoriteResponse, err error) {
	var _args CountUserFavoriteArgs
	_args.Req = Req
	var _result CountUserFavoriteResult
	if err = p.c.Call(ctx, "CountUserFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CountUserTotalFavorited(ctx context.Context, Req *favorite.CountUserTotalFavoritedRequest) (r *favorite.CountUserTotalFavoritedResponse, err error) {
	var _args CountUserTotalFavoritedArgs
	_args.Req = Req
	var _result CountUserTotalFavoritedResult
	if err = p.c.Call(ctx, "CountUserTotalFavorited", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
