package rpcx

import (
	"ddd-template/adapters/rpcx/pb"
	"ddd-template/app"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"net"
)

var ProviderSet = wire.NewSet(NewGrpc, NewDemoGrpcServer)

type GrpcServerImpl struct {
	handlerDemo pb.GreeterServer
	server      *grpc.Server
	opts        []grpc.ServerOption
}

var _ app.GrpcServer = (*GrpcServerImpl)(nil)

//
// Run
// #Description: 执行开启grpc服务
// #receiver g *GrpcServerImpl
// #param addr string
// #return error
func (g *GrpcServerImpl) Run(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	g.server = grpc.NewServer(g.opts...)
	defer g.server.Stop()
	pb.RegisterGreeterServer(g.server, g.handlerDemo)
	if err = g.server.Serve(lis); err != nil {
		return err
	}
	return err
}

//
// SetOptions
// #Description: 设置server插件
// #receiver g *GrpcServerImpl
// #param opts ...grpc.ServerOption
// #return app.GrpcServer
func (g *GrpcServerImpl) SetOptions(opts ...grpc.ServerOption) app.GrpcServer {
	if len(opts) > 0 {
		g.opts = append(g.opts, opts...)
	}
	return g
}

//
// NewGrpc
// #Description: new
// #param opts ...grpc.ServerOption
// #return app.GrpcServer
func NewGrpc(d pb.GreeterServer) app.GrpcServer {
	g := &GrpcServerImpl{}
	g.handlerDemo = d
	return g
}
