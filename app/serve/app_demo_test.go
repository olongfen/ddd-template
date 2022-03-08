package serve

import (
	"context"
	"ddd-template/adapters/repo_mock"
	"ddd-template/domain/entities"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestDemoServerImpl_SayHello(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mockRepo := repo_mock.NewMockDemoInterface(mockCtl)
	logger, _ := zap.NewProduction()

	srv := NewDemoServer(mockRepo, logger)
	data := &entities.Demo{
		Message: "888 " + "hello",
	}
	mockRepo.EXPECT().SayHello(context.Background(), "888").Return(data)
	res, err := srv.SayHello(context.Background(), "888")
	if err != nil {
		assert.Error(t, err)
	}
	logger.Sugar().Info("data:   ", res)
	assert.EqualValues(t, res.Message, data.Message)
}
