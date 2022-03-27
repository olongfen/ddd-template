package restful

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func corsHandler() gin.HandlerFunc {
	corsCfg := cors.DefaultConfig()
	corsCfg.AllowAllOrigins = true
	return cors.New(corsCfg)
}
