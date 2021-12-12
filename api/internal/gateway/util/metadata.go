package util

import (
	"context"
	"errors"
	"strings"

	"github.com/calmato/shs-web/api/pkg/metadata"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	gmd "google.golang.org/grpc/metadata"
)

var errNotExistsAuthorizationHeader = errors.New("util: authorization header is not contain")

func GetAuthToken(c *gin.Context) (string, error) {
	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		return "", errNotExistsAuthorizationHeader
	}

	token := strings.Replace(authorization, "Bearer ", "", 1)
	return token, nil
}

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
