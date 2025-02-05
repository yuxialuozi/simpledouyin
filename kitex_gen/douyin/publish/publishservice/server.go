// Code generated by Kitex v0.12.1. DO NOT EDIT.
package publishservice

import (
	server "github.com/cloudwego/kitex/server"
	publish "simpledouyin/kitex_gen/douyin/publish"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler publish.PublishService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler publish.PublishService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
