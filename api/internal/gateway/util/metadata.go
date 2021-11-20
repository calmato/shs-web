package util

import (
	"context"

	"github.com/calmato/shs-web/api/pkg/metadata"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	gmd "google.golang.org/grpc/metadata"
)

func SetMetadata(c *gin.Context) context.Context {
	ctx := metadata.GinContextToContext(c)

	token := c.GetHeader("Authorization")
	if token != "" {
		ctx = gmd.AppendToOutgoingContext(ctx, "Authorization", token)
	}

	userID := c.GetHeader("userId")
	if userID != "" {
		ctx = gmd.AppendToOutgoingContext(ctx, "userId", userID)
	}

	role := c.GetHeader("role")
	if role != "" {
		ctx = gmd.AppendToOutgoingContext(ctx, "role", role)
	}

	requestID := c.GetHeader("X-Request-ID")
	if requestID == "" {
		requestID = uuid.New().String()
	}
	ctx = gmd.AppendToOutgoingContext(ctx, "X-Request-ID", requestID)

	return ctx
}
