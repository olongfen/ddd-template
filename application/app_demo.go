package application

import (
	"github.com/gin-gonic/gin"
)

// SayHello
// @Tags Demo
// @Summary say hello
// @Description sends a string msg
// @Accept application/json
// @Produce application/json
// @Param msg query string false "message"
// @Router /api/v1/demo  [get]
// @Success 200 {object} Result{code=int,data=string}
func (s *ServerImpl) SayHello(ctx *gin.Context) {
	msg := ctx.Query("msg")
	res := s.demoService.SayHello(ctx.Request.Context(), msg)
	SendSuccess(ctx, res)
	//Fail(ctx, core.Error(errorx.AdminCreateError, errorx.Text(errorx.AdminCreateError)))

}
