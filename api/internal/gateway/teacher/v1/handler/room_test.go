package handler

import (
	"context"
	"net/http"
	"testing"

	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/request"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/golang/mock/gomock"
)

func TestGetRoomsTotal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &classroom.GetRoomsTotalRequest{}
				out := &classroom.GetRoomsTotalResponse{Total: 3}
				mocks.classroom.EXPECT().GetRoomsTotal(gomock.Any(), in).Return(out, nil)
			},
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.RoomsTotalResponse{
					Total: 3,
				},
			},
		},
		{
			name: "failed to get rooms total",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &classroom.GetRoomsTotalRequest{}
				mocks.classroom.EXPECT().GetRoomsTotal(gomock.Any(), in).Return(nil, errmock)
			},
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/rooms"
			req := newHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestUpdateRoomsTotal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		req    *request.UpdateRoomsTotalRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &classroom.UpdateRoomsTotalRequest{Total: 3}
				out := &classroom.UpdateRoomsTotalResponse{}
				mocks.classroom.EXPECT().UpdateRoomsTotal(gomock.Any(), in).Return(out, nil)
			},
			req: &request.UpdateRoomsTotalRequest{
				Total: 3,
			},
			expect: &testResponse{
				code: http.StatusNoContent,
			},
		},
		{
			name: "failed to update rooms total",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &classroom.UpdateRoomsTotalRequest{Total: 3}
				mocks.classroom.EXPECT().UpdateRoomsTotal(gomock.Any(), in).Return(nil, errmock)
			},
			req: &request.UpdateRoomsTotalRequest{
				Total: 3,
			},
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/rooms"
			req := newHTTPRequest(t, http.MethodPatch, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}
