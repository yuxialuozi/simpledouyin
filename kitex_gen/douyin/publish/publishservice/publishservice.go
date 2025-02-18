// Code generated by Kitex v0.12.1. DO NOT EDIT.

package publishservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	publish "simpledouyin/kitex_gen/douyin/publish"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"CreateVideo": kitex.NewMethodInfo(
		createVideoHandler,
		newCreateVideoArgs,
		newCreateVideoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ListVideo": kitex.NewMethodInfo(
		listVideoHandler,
		newListVideoArgs,
		newListVideoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"CountVideo": kitex.NewMethodInfo(
		countVideoHandler,
		newCountVideoArgs,
		newCountVideoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	publishServiceServiceInfo                = NewServiceInfo()
	publishServiceServiceInfoForClient       = NewServiceInfoForClient()
	publishServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return publishServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return publishServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return publishServiceServiceInfoForClient
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
	serviceName := "PublishService"
	handlerType := (*publish.PublishService)(nil)
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
		"PackageName": "simpledouyin.publish",
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

func createVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(publish.CreateVideoRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(publish.PublishService).CreateVideo(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CreateVideoArgs:
		success, err := handler.(publish.PublishService).CreateVideo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateVideoResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCreateVideoArgs() interface{} {
	return &CreateVideoArgs{}
}

func newCreateVideoResult() interface{} {
	return &CreateVideoResult{}
}

type CreateVideoArgs struct {
	Req *publish.CreateVideoRequest
}

func (p *CreateVideoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(publish.CreateVideoRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateVideoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateVideoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateVideoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CreateVideoArgs) Unmarshal(in []byte) error {
	msg := new(publish.CreateVideoRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateVideoArgs_Req_DEFAULT *publish.CreateVideoRequest

func (p *CreateVideoArgs) GetReq() *publish.CreateVideoRequest {
	if !p.IsSetReq() {
		return CreateVideoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateVideoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CreateVideoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CreateVideoResult struct {
	Success *publish.CreateVideoResponse
}

var CreateVideoResult_Success_DEFAULT *publish.CreateVideoResponse

func (p *CreateVideoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(publish.CreateVideoResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateVideoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateVideoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateVideoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CreateVideoResult) Unmarshal(in []byte) error {
	msg := new(publish.CreateVideoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateVideoResult) GetSuccess() *publish.CreateVideoResponse {
	if !p.IsSetSuccess() {
		return CreateVideoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateVideoResult) SetSuccess(x interface{}) {
	p.Success = x.(*publish.CreateVideoResponse)
}

func (p *CreateVideoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateVideoResult) GetResult() interface{} {
	return p.Success
}

func listVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(publish.ListVideoRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(publish.PublishService).ListVideo(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ListVideoArgs:
		success, err := handler.(publish.PublishService).ListVideo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListVideoResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newListVideoArgs() interface{} {
	return &ListVideoArgs{}
}

func newListVideoResult() interface{} {
	return &ListVideoResult{}
}

type ListVideoArgs struct {
	Req *publish.ListVideoRequest
}

func (p *ListVideoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(publish.ListVideoRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListVideoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListVideoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListVideoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ListVideoArgs) Unmarshal(in []byte) error {
	msg := new(publish.ListVideoRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListVideoArgs_Req_DEFAULT *publish.ListVideoRequest

func (p *ListVideoArgs) GetReq() *publish.ListVideoRequest {
	if !p.IsSetReq() {
		return ListVideoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListVideoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ListVideoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ListVideoResult struct {
	Success *publish.ListVideoResponse
}

var ListVideoResult_Success_DEFAULT *publish.ListVideoResponse

func (p *ListVideoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(publish.ListVideoResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListVideoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListVideoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListVideoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ListVideoResult) Unmarshal(in []byte) error {
	msg := new(publish.ListVideoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListVideoResult) GetSuccess() *publish.ListVideoResponse {
	if !p.IsSetSuccess() {
		return ListVideoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListVideoResult) SetSuccess(x interface{}) {
	p.Success = x.(*publish.ListVideoResponse)
}

func (p *ListVideoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListVideoResult) GetResult() interface{} {
	return p.Success
}

func countVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(publish.CountVideoRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(publish.PublishService).CountVideo(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CountVideoArgs:
		success, err := handler.(publish.PublishService).CountVideo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CountVideoResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCountVideoArgs() interface{} {
	return &CountVideoArgs{}
}

func newCountVideoResult() interface{} {
	return &CountVideoResult{}
}

type CountVideoArgs struct {
	Req *publish.CountVideoRequest
}

func (p *CountVideoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(publish.CountVideoRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CountVideoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CountVideoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CountVideoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CountVideoArgs) Unmarshal(in []byte) error {
	msg := new(publish.CountVideoRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CountVideoArgs_Req_DEFAULT *publish.CountVideoRequest

func (p *CountVideoArgs) GetReq() *publish.CountVideoRequest {
	if !p.IsSetReq() {
		return CountVideoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CountVideoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CountVideoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CountVideoResult struct {
	Success *publish.CountVideoResponse
}

var CountVideoResult_Success_DEFAULT *publish.CountVideoResponse

func (p *CountVideoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(publish.CountVideoResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CountVideoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CountVideoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CountVideoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CountVideoResult) Unmarshal(in []byte) error {
	msg := new(publish.CountVideoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CountVideoResult) GetSuccess() *publish.CountVideoResponse {
	if !p.IsSetSuccess() {
		return CountVideoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CountVideoResult) SetSuccess(x interface{}) {
	p.Success = x.(*publish.CountVideoResponse)
}

func (p *CountVideoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CountVideoResult) GetResult() interface{} {
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

func (p *kClient) CreateVideo(ctx context.Context, Req *publish.CreateVideoRequest) (r *publish.CreateVideoResponse, err error) {
	var _args CreateVideoArgs
	_args.Req = Req
	var _result CreateVideoResult
	if err = p.c.Call(ctx, "CreateVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListVideo(ctx context.Context, Req *publish.ListVideoRequest) (r *publish.ListVideoResponse, err error) {
	var _args ListVideoArgs
	_args.Req = Req
	var _result ListVideoResult
	if err = p.c.Call(ctx, "ListVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CountVideo(ctx context.Context, Req *publish.CountVideoRequest) (r *publish.CountVideoResponse, err error) {
	var _args CountVideoArgs
	_args.Req = Req
	var _result CountVideoResult
	if err = p.c.Call(ctx, "CountVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
