package util

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	gmd "google.golang.org/grpc/metadata"
)

func TestGetAuthToken(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		ctx    *gin.Context
		expect string
		isErr  bool
	}{
		{
			name: "success",
			ctx: func() *gin.Context {
				gin.SetMode(gin.TestMode)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = &http.Request{Header: http.Header{}}
				c.Request.Header.Set("Authorization", "Bearer xxxxxx")
				return c
			}(),
			expect: "xxxxxx",
			isErr:  false,
		},
		{
			name: "not exists authorization header",
			ctx: func() *gin.Context {
				gin.SetMode(gin.TestMode)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = &http.Request{Header: http.Header{}}
				return c
			}(),
			expect: "",
			isErr:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual, err := GetAuthToken(tt.ctx)
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestSetMetadata(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{Header: http.Header{}}
	c.Request.Header.Set("Authorization", "Bearer xxxxxx")
	c.Request.Header.Set("userId", "userid")
	c.Request.Header.Set("role", "role")

	ctx := context.Background()
	ctx = context.WithValue(ctx, "Authorization", "Bearer xxxxxx")
	ctx = context.WithValue(ctx, "userId", "userid")
	ctx = context.WithValue(ctx, "role", "role")

	actual := SetMetadata(c)
	assert.IsType(t, ctx, actual)
	md, ok := gmd.FromOutgoingContext(actual)
	if ok {
		assert.Contains(t, md.Get("Authorization"), "Bearer xxxxxx")
		assert.Contains(t, md.Get("userId"), "userid")
		assert.Contains(t, md.Get("role"), "role")
		assert.NotContains(t, md.Get("X-Request-ID"), "")
	} else {
		t.FailNow()
	}
}
