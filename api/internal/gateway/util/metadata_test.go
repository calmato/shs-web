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
