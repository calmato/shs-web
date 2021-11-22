package util

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorResponse struct {
	Status  int    `json:"status"`  // ステータスコード
	Message string `json:"message"` // エラー概要
	Detail  string `json:"detail"`  // エラー詳細
}

func NewErrorResponse(err error) (*ErrorResponse, int) {
	res := &ErrorResponse{}

	switch status.Code(err) {
	case codes.Canceled:
		res.Status = 499 // client closed request
	case codes.Unknown, codes.Internal, codes.DataLoss:
		res.Status = http.StatusInternalServerError
	case codes.InvalidArgument, codes.ResourceExhausted, codes.OutOfRange:
		res.Status = http.StatusBadRequest
	case codes.DeadlineExceeded:
		res.Status = http.StatusGatewayTimeout
	case codes.NotFound:
		res.Status = http.StatusNotFound
	case codes.AlreadyExists:
		res.Status = http.StatusConflict
	case codes.PermissionDenied:
		res.Status = http.StatusForbidden
	case codes.FailedPrecondition:
		res.Status = http.StatusPreconditionFailed
	case codes.Aborted:
		res.Status = http.StatusConflict
	case codes.Unimplemented:
		res.Status = http.StatusNotImplemented
	case codes.Unavailable:
		res.Status = http.StatusBadGateway
	case codes.Unauthenticated:
		res.Status = http.StatusUnauthorized
	default:
		res.Status = http.StatusInternalServerError
		res.Message = "unknown error code"
		res.Detail = fmt.Sprintf("util: %v", err)
		return res, http.StatusInternalServerError
	}

	res.Message = http.StatusText(res.Status)
	res.Detail = err.Error()
	return res, res.Status
}
