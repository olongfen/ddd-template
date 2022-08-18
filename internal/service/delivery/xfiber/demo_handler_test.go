package xfiber

import (
	"context"
	"ddd-template/internal/common/utils"
	"ddd-template/internal/domain"
	mock_domain "ddd-template/internal/infra/mock/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
	"time"
)

func TestDemoHandler_Handler(t *testing.T) {
	app := fiber.New()
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	uc := mock_domain.NewMockIDemoUseCase(ctl)
	d := NewDemoHandler(uc)
	d.Handler(app)
	// get
	getDemo := &domain.Demo{
		Model:   utils.Model{ID: 1, CreatedAt: utils.JSONTime{Time: time.Now()}},
		Message: "wwwwwwwwwwwwwwwwwww",
	}
	uc.EXPECT().Get(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, id uint) (*domain.Demo, error) {
		return getDemo, nil
	})
	resp, err := app.Test(httptest.NewRequest("GET", "/demo/1", nil))
	assert.Equal(t, nil, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var (
		getDemoResp = new(domain.Demo)
		respMap     = make(map[string]interface{})
	)
	body, err := io.ReadAll(resp.Body)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, jsoniter.Unmarshal(body, &respMap))
	assert.Equal(t, nil, utils.Copier(respMap["data"], getDemoResp))
	assert.Equal(t, getDemo.Message, getDemoResp.Message)
	resp.Body.Close()

}
