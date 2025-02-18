// Code generated by Kitex v0.12.1. DO NOT EDIT.

package commentservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	comment "simpledouyin/kitex_gen/douyin/comment"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"ActionComment": kitex.NewMethodInfo(
		actionCommentHandler,
		newActionCommentArgs,
		newActionCommentResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ListComment": kitex.NewMethodInfo(
		listCommentHandler,
		newListCommentArgs,
		newListCommentResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"CountComment": kitex.NewMethodInfo(
		countCommentHandler,
		newCountCommentArgs,
		newCountCommentResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	commentServiceServiceInfo                = NewServiceInfo()
	commentServiceServiceInfoForClient       = NewServiceInfoForClient()
	commentServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return commentServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return commentServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return commentServiceServiceInfoForClient
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
	serviceName := "CommentService"
	handlerType := (*comment.CommentService)(nil)
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
		"PackageName": "simpledouyin.comment",
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

func actionCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(comment.ActionCommentRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(comment.CommentService).ActionComment(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ActionCommentArgs:
		success, err := handler.(comment.CommentService).ActionComment(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ActionCommentResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newActionCommentArgs() interface{} {
	return &ActionCommentArgs{}
}

func newActionCommentResult() interface{} {
	return &ActionCommentResult{}
}

type ActionCommentArgs struct {
	Req *comment.ActionCommentRequest
}

func (p *ActionCommentArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(comment.ActionCommentRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ActionCommentArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ActionCommentArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ActionCommentArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ActionCommentArgs) Unmarshal(in []byte) error {
	msg := new(comment.ActionCommentRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ActionCommentArgs_Req_DEFAULT *comment.ActionCommentRequest

func (p *ActionCommentArgs) GetReq() *comment.ActionCommentRequest {
	if !p.IsSetReq() {
		return ActionCommentArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ActionCommentArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ActionCommentArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ActionCommentResult struct {
	Success *comment.ActionCommentResponse
}

var ActionCommentResult_Success_DEFAULT *comment.ActionCommentResponse

func (p *ActionCommentResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(comment.ActionCommentResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ActionCommentResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ActionCommentResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ActionCommentResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ActionCommentResult) Unmarshal(in []byte) error {
	msg := new(comment.ActionCommentResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ActionCommentResult) GetSuccess() *comment.ActionCommentResponse {
	if !p.IsSetSuccess() {
		return ActionCommentResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ActionCommentResult) SetSuccess(x interface{}) {
	p.Success = x.(*comment.ActionCommentResponse)
}

func (p *ActionCommentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ActionCommentResult) GetResult() interface{} {
	return p.Success
}

func listCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(comment.ListCommentRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(comment.CommentService).ListComment(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ListCommentArgs:
		success, err := handler.(comment.CommentService).ListComment(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListCommentResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newListCommentArgs() interface{} {
	return &ListCommentArgs{}
}

func newListCommentResult() interface{} {
	return &ListCommentResult{}
}

type ListCommentArgs struct {
	Req *comment.ListCommentRequest
}

func (p *ListCommentArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(comment.ListCommentRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListCommentArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListCommentArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListCommentArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ListCommentArgs) Unmarshal(in []byte) error {
	msg := new(comment.ListCommentRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListCommentArgs_Req_DEFAULT *comment.ListCommentRequest

func (p *ListCommentArgs) GetReq() *comment.ListCommentRequest {
	if !p.IsSetReq() {
		return ListCommentArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListCommentArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ListCommentArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ListCommentResult struct {
	Success *comment.ListCommentResponse
}

var ListCommentResult_Success_DEFAULT *comment.ListCommentResponse

func (p *ListCommentResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(comment.ListCommentResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListCommentResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListCommentResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListCommentResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ListCommentResult) Unmarshal(in []byte) error {
	msg := new(comment.ListCommentResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListCommentResult) GetSuccess() *comment.ListCommentResponse {
	if !p.IsSetSuccess() {
		return ListCommentResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListCommentResult) SetSuccess(x interface{}) {
	p.Success = x.(*comment.ListCommentResponse)
}

func (p *ListCommentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListCommentResult) GetResult() interface{} {
	return p.Success
}

func countCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(comment.CountCommentRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(comment.CommentService).CountComment(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CountCommentArgs:
		success, err := handler.(comment.CommentService).CountComment(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CountCommentResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCountCommentArgs() interface{} {
	return &CountCommentArgs{}
}

func newCountCommentResult() interface{} {
	return &CountCommentResult{}
}

type CountCommentArgs struct {
	Req *comment.CountCommentRequest
}

func (p *CountCommentArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(comment.CountCommentRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CountCommentArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CountCommentArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CountCommentArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CountCommentArgs) Unmarshal(in []byte) error {
	msg := new(comment.CountCommentRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CountCommentArgs_Req_DEFAULT *comment.CountCommentRequest

func (p *CountCommentArgs) GetReq() *comment.CountCommentRequest {
	if !p.IsSetReq() {
		return CountCommentArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CountCommentArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CountCommentArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CountCommentResult struct {
	Success *comment.CountCommentResponse
}

var CountCommentResult_Success_DEFAULT *comment.CountCommentResponse

func (p *CountCommentResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(comment.CountCommentResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CountCommentResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CountCommentResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CountCommentResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CountCommentResult) Unmarshal(in []byte) error {
	msg := new(comment.CountCommentResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CountCommentResult) GetSuccess() *comment.CountCommentResponse {
	if !p.IsSetSuccess() {
		return CountCommentResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CountCommentResult) SetSuccess(x interface{}) {
	p.Success = x.(*comment.CountCommentResponse)
}

func (p *CountCommentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CountCommentResult) GetResult() interface{} {
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

func (p *kClient) ActionComment(ctx context.Context, Req *comment.ActionCommentRequest) (r *comment.ActionCommentResponse, err error) {
	var _args ActionCommentArgs
	_args.Req = Req
	var _result ActionCommentResult
	if err = p.c.Call(ctx, "ActionComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListComment(ctx context.Context, Req *comment.ListCommentRequest) (r *comment.ListCommentResponse, err error) {
	var _args ListCommentArgs
	_args.Req = Req
	var _result ListCommentResult
	if err = p.c.Call(ctx, "ListComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CountComment(ctx context.Context, Req *comment.CountCommentRequest) (r *comment.CountCommentResponse, err error) {
	var _args CountCommentArgs
	_args.Req = Req
	var _result CountCommentResult
	if err = p.c.Call(ctx, "CountComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
