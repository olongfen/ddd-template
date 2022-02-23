package application

import (
	"context"
	repo_mock "ddd-template/adapters/mock/repo"
	"ddd-template/domain/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerImpl_SayHello(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mockRepo := repo_mock.NewMockDemoInterface(mockCtl)
	logger, _ := zap.NewProduction()
	demo := service.NewDemoService(mockRepo, logger)
	app := NewServer(demo, logger)
	w := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(w)
	uri := "/api/v1/test/demo"
	engine.GET(uri, app.SayHello)
	mockRepo.EXPECT().SayHello(context.Background(), "hello").Return("hello" + " world")
	req, _ := http.NewRequest("GET", uri+"?msg=hello", nil)
	engine.ServeHTTP(w, req)
	assert.EqualValues(t, w.Code, 200)
	res := gin.H{}
	json.NewDecoder(w.Body).Decode(&res)
	assert.EqualValues(t, res["code"], 0)
	logger.Sugar().Info(res)
	assert.EqualValues(t, res["data"], "hello"+" world")

}
