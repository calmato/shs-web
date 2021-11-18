package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func newRouter(reg *registry, opts ...gin.HandlerFunc) *gin.Engine {
	rt := gin.Default()
	rt.Use(opts...)

	// other routes
	rt.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})

	rt.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, "not found")
	})

	return rt
}
