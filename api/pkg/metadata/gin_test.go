package metadata

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGinContextToContext(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		ctx    *gin.Context
		expect bool
	}{
		{
			name: "success",
			ctx: func() *gin.Context {
				gin.SetMode(gin.TestMode)
				ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
				ctx.Request = httptest.NewRequest(http.MethodGet, "http://example.com", nil)
				return ctx
			}(),
			expect: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := GinContextToContext(tt.ctx)
			assert.Equal(t, tt.expect, ctx != nil)
		})
	}
}

func TestGinContextFromContext(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = httptest.NewRequest(http.MethodGet, "http://example.com", nil)

	tests := []struct {
		name   string
		ctx    context.Context
		expect *gin.Context
		isErr  bool
	}{
		{
			name:   "success",
			ctx:    context.WithValue(ctx.Request.Context(), GinContext, ctx),
			expect: ctx,
			isErr:  false,
		},
		{
			name:   "not retrieve context",
			ctx:    context.WithValue(ctx.Request.Context(), "test", ctx),
			expect: nil,
			isErr:  true,
		},
		{
			name:   "invalid context",
			ctx:    context.WithValue(ctx.Request.Context(), GinContext, context.Background()),
			expect: nil,
			isErr:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := GinContextFromContext(tt.ctx)
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
