package rpcx

import (
	"context"
	"ddd-template/adapters/rpcx/pb"
	"ddd-template/app/serve"
	"ddd-template/common/schema"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DemoGrpcServer struct {
	pb.UnimplementedGreeterServer
	log    *zap.Logger
	server serve.DemoServer
}

//
// NewDemoGrpcServer
// #Description: new
// #param srv serve.DemoServer
// #param logger *zap.Logger
// #return pb.GreeterServer
func NewDemoGrpcServer(srv serve.DemoServer, logger *zap.Logger) pb.GreeterServer {
	return &DemoGrpcServer{server: srv, log: logger}
}

func demoSchema2pb(in *schema.DemoInfo) (out *pb.DemoInfo) {
	out = new(pb.DemoInfo)
	out.Message = in.Message
	out.Id = int32(in.ID)
	out.CreatedAt = timestamppb.New(in.CreatedAt)
	out.UpdatedAt = timestamppb.New(in.UpdatedAt)
	return
}

//
// SayHello
// #Description: say hello
// #receiver s *DemoGrpcServer
// #param ctx context.Context
// #param in *pb.HelloRequest
// #return res *pb.DemoInfo
// #return err error
func (s *DemoGrpcServer) SayHello(ctx context.Context, in *pb.HelloRequest) (res *pb.DemoInfo, err error) {
	var (
		data *schema.DemoInfo
	)
	if data, err = s.server.SayHello(ctx, in.Msg); err != nil {
		return
	}
	res = demoSchema2pb(data)
	return
}
