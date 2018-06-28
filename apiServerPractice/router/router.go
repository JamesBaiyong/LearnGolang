package router

import (
	"github.com/gin-gonic/gin"
	"LearnGolang/apiServerPractice/handler/sd"
	"LearnGolang/apiServerPractice/router/middleware"
	"net/http"
)

func Load(router *gin.Engine, mw ...gin.HandlerFunc)  *gin.Engine{
	//在处理某些请求时可能因为程序 bug 或者其他异常情况导致程序 panic，这时候为了不影响下一次请求的调用，需要通过 gin.Recovery()来恢复 API 服务器
	router.Use(gin.Recovery())
	router.Use(middleware.NoCache)
	router.Use(middleware.Options)
	router.Use(middleware.Secure)
	router.Use(mw...)

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound,"Error API path.")
	})

	r1 := router.Group("/sd")
	{
		r1.GET("/health",sd.HealthCheck)
		r1.GET("/disk",sd.DiskCheck)
		r1.GET("/cpu",sd.CPUCheck)
		r1.GET("/ram",sd.RAMCheck)

	}
	return router
}
