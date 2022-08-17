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
	ctx := context.Background()
	mockRepo.EXPECT().Get(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, d *domain.Demo) error {
		*d = *data
		return nil
	})
	var (
		res *domain.Demo
		err error
	)
	if res, err = srv.Get(ctx, 1); err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, *data, *res)
}
