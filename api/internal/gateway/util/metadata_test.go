package util

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

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
	assert.Equal(t, ctx.Value("Authorization").(string), c.GetHeader("Authorization"))
	assert.Equal(t, ctx.Value("userId").(string), c.GetHeader("userId"))
	assert.Equal(t, ctx.Value("role").(string), c.GetHeader("role"))
	assert.NotZero(t, c.GetHeader("X-Request-ID"))
}
