package rpcx

import (
	"ddd-template/api/v1"
	"ddd-template/internal/app"
	"ddd-template/internal/common/conf"
	"ddd-template/internal/common/xlog"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"net"
)

type GrpcServer struct {
	handlerDemo v1.GreeterServer
	server      *grpc.Server
	opts        []grpc.ServerOption
	cfg         conf.GRpc
}

var _ app.GrpcServer = (*GrpcServer)(nil)

//
// Run
// #Description: 执行开启grpc服务
// #receiver g *GrpcServer
// #param addr string
// #return error
func (g *GrpcServer) Run() error {
	addr := fmt.Sprintf("%s:%d", g.cfg.Host, g.cfg.Port)
	xlog.Log.Sugar().Infof("grpc server run in: %s", addr)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	g.server = grpc.NewServer(g.opts...)
	defer g.server.Stop()
	v1.RegisterGreeterServer(g.server, g.handlerDemo)
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
func NewGrpc(d v1.GreeterServer, cfg *conf.Configs) app.GrpcServer {
	g := &GrpcServer{}
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
