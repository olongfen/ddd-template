package usecase

import (
	"context"
	"ddd-template/internal/common/utils"
	mock_domain "ddd-template/internal/infra/mock/domain"

	"ddd-template/internal/domain"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestDemoServer_SayHello(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mockRepo := mock_domain.NewMockIDemoRepo(mockCtl)
	logger, _ := zap.NewDevelopment()
	tx := mock_domain.NewMockITransaction(mockCtl)
	srv := NewDemoServer(mockRepo, tx, logger)
	data := &domain.Demo{
		Model: utils.Model{
			ID:        1,
			CreatedAt: utils.JSONTime{Time: time.Now()},
			UpdatedAt: utils.JSONTime{Time: time.Now()},
		},
		Message: "888 " + "hello",
	}
	mockRepo.EXPECT().SayHello(context.Background(), "1818").Return(data)
	var (
		res *domain.Demo
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
