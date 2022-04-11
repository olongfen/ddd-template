package rpcx

import (
	"ddd-template/adapters/rpcx/pb"
	"ddd-template/app"
	"ddd-template/common/conf"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
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
// NewGrpc
// #Description: new
// #param opts ...grpc.ServerOption
// #return app.GrpcServer
func NewGrpc(d pb.GreeterServer, cfg conf.Configs) app.GrpcServer {
	g := &GrpcServerImpl{}
	g.handlerDemo = d
	if cfg.Server.GRpc.TLS {
		creds, err := credentials.NewServerTLSFromFile(cfg.Server.GRpc.PEMFile, cfg.Server.GRpc.KeyFile)
		if err != nil {
			grpclog.Fatalf("Failed to generate credentials %v", err)
		}
		g.opts = append(g.opts, grpc.Creds(creds))

	}
	return g
}
