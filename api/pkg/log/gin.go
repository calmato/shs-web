package log

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func NewGinMiddleware(params *Params) (gin.HandlerFunc, error) {
	logger, err := NewLogger(params)
	if err != nil {
		return nil, err
	}
	return ginzap.Ginzap(logger, time.RFC3339, true), nil
}
