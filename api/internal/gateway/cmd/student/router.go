package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func newRouter(reg *registry, opts ...gin.HandlerFunc) *gin.Engine {
	rt := gin.Default()
	rt.Use(opts...)

	// auth required routes
	authV1Group := rt.Group("")
	authV1Group.Use(reg.v1.Authentication())
	reg.v1.AuthRoutes(authV1Group)

	// public routes
	reg.v1.NoAuthRoutes(rt.Group(""))

	// other routes
	rt.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})

	rt.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, "not found")
	})

	return rt
}
