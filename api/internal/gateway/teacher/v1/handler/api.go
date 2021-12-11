package handler

import (
	"sync"
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/**
 * ###############################################
 * handler
 * ###############################################
 */
type APIV1Handler interface {
	AuthRoutes(rg *gin.RouterGroup)   // 認証済みでアクセス可能なエンドポイント一覧
	AdminRoutes(rg *gin.RouterGroup)  // 管理者のみアクセス可能なエンドポイント一覧
	NoAuthRoutes(rg *gin.RouterGroup) // 未認証でもアクセス可能なエンドポイント一覧
}

type apiV1Handler struct {
	now         func() time.Time
	logger      *zap.Logger
	sharedGroup *singleflight.Group
	waitGroup   *sync.WaitGroup
	user        user.UserServiceClient
}

type Params struct {
	UserService user.UserServiceClient
	Logger      *zap.Logger
	WaitGroup   *sync.WaitGroup
}

func NewAPIV1Handler(params *Params) APIV1Handler {
	return &apiV1Handler{
		now:         jst.Now,
		logger:      params.Logger,
		waitGroup:   params.WaitGroup,
		user:        params.UserService,
		sharedGroup: &singleflight.Group{},
	}
}

/**
 * ###############################################
 * routes
 * ###############################################
 */
func (h *apiV1Handler) AuthRoutes(rg *gin.RouterGroup) {}

func (h *apiV1Handler) AdminRoutes(rg *gin.RouterGroup) {
	rg.POST("/v1/teachers", h.CreateTeacher)
}

func (h *apiV1Handler) NoAuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/v1/hello", h.Hello)
}

/**
 * ###############################################
 * error handling
 * ###############################################
 */
func httpError(ctx *gin.Context, err error) {
	res, code := util.NewErrorResponse(err)
	ctx.JSON(code, res)
	ctx.Abort()
}

func badRequest(ctx *gin.Context, err error) {
	httpError(ctx, status.Error(codes.InvalidArgument, err.Error()))
}

/**
 * ###############################################
 * other
 * ###############################################
 */
