package serve

import (
	"context"
	"ddd-template/adapters/repo_mock"
	"ddd-template/app/dto"
	"ddd-template/domain/entities"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestDemoServerImpl_SayHello(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mockRepo := repo_mock.NewMockDemoInterface(mockCtl)
	logger, _ := zap.NewDevelopment()
	srv := NewDemoServer(mockRepo, logger)
	data := &entities.Demo{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Message: "888 " + "hello",
	}
	mockRepo.EXPECT().SayHello(context.Background(), "1818").Return(data)
	var (
		res *dto.DemoInfo
		err error
	)
	if res, err = srv.SayHello(context.Background(), "1818"); err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, data.Message, res.Message)
	assert.EqualValues(t, data.ID, res.ID)
	assert.EqualValues(t, data.UpdatedAt, res.UpdatedAt)
	assert.EqualValues(t, data.CreatedAt, res.CreatedAt)
}
