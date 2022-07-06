package service

import (
	"context"
	v1 "ddd-template/api/v1"
	"ddd-template/internal/domain"
	"go.uber.org/zap"
)

type DemoService struct {
	usecase domain.IDemoUsecase
	log     *zap.Logger
	v1.UnimplementedGreeterServer
}

func NewDemoService(usecase domain.IDemoUsecase, logger *zap.Logger) v1.GreeterServer {
	return &DemoService{usecase: usecase, log: logger}
}

// SayHello
// @Tags Demo
// @Summary
// @Description
// @Accept application/json
// @Produce application/json
// @Param msg query string false "message"
// @Router /api/v1/hello [get]
// @Success 200 {object} response.HTTPServerResponse{code=int,data=v1.DemoInfo}
// @Failure 500 {object}    interface{}  "服务器内部错误"
func (d *DemoService) SayHello(ctx context.Context, req *v1.HelloRequest) (ret *v1.DemoInfo, err error) {
	var (
		data *domain.Demo
	)
	if data, err = d.usecase.SayHello(ctx, req.Msg); err != nil {
		return
	}
	ret = new(v1.DemoInfo)
	ret.Message = data.Message + "<<<action:>>>" + req.Action
	return
}
