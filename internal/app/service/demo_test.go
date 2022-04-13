package service

import (
	"context"
	v1 "ddd-template/api/v1"
	mock_domain "ddd-template/internal/adapters/mock/domain"
	"ddd-template/internal/domain"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestDemoService_SayHello(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	usecase := mock_domain.NewMockIDemoUsecase(mockCtl)
	logger, _ := zap.NewDevelopment()
	srv := NewDemoService(usecase, logger)
	data := &domain.Demo{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Message: "888 " + "hello",
	}
	usecase.EXPECT().SayHello(context.Background(), "1818").Return(data, nil)
	var (
		res *v1.DemoInfo
		err error
	)
	if res, err = srv.SayHello(context.Background(), &v1.HelloRequest{Msg: "1818"}); err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, data.Message, res.Message)
}
