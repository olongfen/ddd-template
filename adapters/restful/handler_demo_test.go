package restful

import (
	"context"
	"ddd-template/adapters/repo_mock"
	"ddd-template/app/serve"
	"ddd-template/domain/entities"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDemoCtl_SayHello(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mockRepo := repo_mock.NewMockDemoInterface(mockCtl)
	logger, _ := zap.NewProduction()

	w := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(w)
	uri := "/demo"
	ctl := NewDemoCtl(serve.NewDemoServer(mockRepo, logger), logger)
	engine.GET(uri, ctl.SayHello)
	data := &entities.Demo{
		Message: "888 " + "hello",
	}
	mockRepo.EXPECT().SayHello(context.Background(), "888").Return(data)
	req, _ := http.NewRequest("GET", uri+"?msg=888", nil)
	engine.ServeHTTP(w, req)
	assert.EqualValues(t, w.Code, 200)
	res := gin.H{}
	json.NewDecoder(w.Body).Decode(&res)
	logger.Sugar().Info("data:   ", res)
	assert.EqualValues(t, res["code"], 0)
	assert.EqualValues(t, res["data"].(map[string]interface{})["message"], "888 hello")

}
