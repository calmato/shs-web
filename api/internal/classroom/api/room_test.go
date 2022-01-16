package api

import (
	"context"
	"testing"

	"github.com/calmato/shs-web/api/internal/classroom/entity"
	"github.com/calmato/shs-web/api/internal/classroom/validation"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func TestGetRoom(t *testing.T) {
	t.Parallel()

	now := jst.Now()
	req := &classroom.GetRoomRequest{
		Id: 1,
	}
	room := &entity.Room{
		ID:        1,
		CreatedAt: now,
		UpdatedAt: now,
	}
	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *classroom.GetRoomRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().GetRoom(req).Return(nil)
				mocks.db.Room.EXPECT().Get(ctx, int32(1)).Return(room, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &classroom.GetRoomResponse{
					Room: &classroom.Room{
						Id:        1,
						CreatedAt: now.Unix(),
						UpdatedAt: now.Unix(),
					},
				},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &classroom.GetRoomRequest{}
				mocks.validator.EXPECT().GetRoom(req).Return(validation.ErrRequestValidation)
			},
			req: &classroom.GetRoomRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to get room",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().GetRoom(req).Return(nil)
				mocks.db.Room.EXPECT().Get(ctx, int32(1)).Return(nil, errmock)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, testGRPC(tt.setup, tt.expect, func(ctx context.Context, service *classroomService) (proto.Message, error) {
			return service.GetRoom(ctx, tt.req)
		}))
	}
}

func TestGetRoomsTotal(t *testing.T) {
	t.Parallel()

	req := &classroom.GetRoomsTotalRequest{}
	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *classroom.GetRoomsTotalRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().GetRoomsTotal(req).Return(nil)
				mocks.db.Room.EXPECT().Count(ctx).Return(int64(3), nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &classroom.GetRoomsTotalResponse{
					Total: 3,
				},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &classroom.GetRoomsTotalRequest{}
				mocks.validator.EXPECT().GetRoomsTotal(req).Return(validation.ErrRequestValidation)
			},
			req: &classroom.GetRoomsTotalRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to get rooms total",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().GetRoomsTotal(req).Return(nil)
				mocks.db.Room.EXPECT().Count(ctx).Return(int64(0), errmock)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, testGRPC(tt.setup, tt.expect, func(ctx context.Context, service *classroomService) (proto.Message, error) {
			return service.GetRoomsTotal(ctx, tt.req)
		}))
	}
}

func TestUpdateRoomsTotal(t *testing.T) {
	t.Parallel()

	req := &classroom.UpdateRoomsTotalRequest{
		Total: 3,
	}
	rooms := entity.Rooms{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	}
	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *classroom.UpdateRoomsTotalRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().UpdateRoomsTotal(req).Return(nil)
				mocks.db.Room.EXPECT().Replace(ctx, rooms).Return(nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &classroom.UpdateRoomsTotalResponse{},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &classroom.UpdateRoomsTotalRequest{}
				mocks.validator.EXPECT().UpdateRoomsTotal(req).Return(validation.ErrRequestValidation)
			},
			req: &classroom.UpdateRoomsTotalRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().UpdateRoomsTotal(req).Return(nil)
				mocks.db.Room.EXPECT().Replace(ctx, rooms).Return(errmock)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, testGRPC(tt.setup, tt.expect, func(ctx context.Context, service *classroomService) (proto.Message, error) {
			return service.UpdateRoomsTotal(ctx, tt.req)
		}))
	}
}
