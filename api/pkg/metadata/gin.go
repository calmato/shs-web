package metadata

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	errInvalidContext     = errors.New("metadata: this context has wrong type")
	errNotRetrieveContext = errors.New("metadata: this context could not retrieve")
)

type key int

// GinContext - Context変換用のキー
var GinContext key

// GinContextToContext - gin.Context -> context.Context
func GinContextToContext(ctx *gin.Context) context.Context {
	return context.WithValue(ctx.Request.Context(), GinContext, ctx)
}

// GinContextFromContext - context.Context -> gin.Context
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(GinContext)
	if ginContext == nil {
		return nil, errNotRetrieveContext
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		return nil, errInvalidContext
	}

	return gc, nil
}
