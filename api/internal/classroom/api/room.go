package api

import (
	"context"

	"github.com/calmato/shs-web/api/internal/classroom/entity"
	"github.com/calmato/shs-web/api/proto/classroom"
)

func (s *classroomService) GetRoom(
	ctx context.Context, req *classroom.GetRoomRequest,
) (*classroom.GetRoomResponse, error) {
	if err := s.validator.GetRoom(req); err != nil {
		return nil, gRPCError(err)
	}

	room, err := s.db.Room.Get(ctx, req.Id)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &classroom.GetRoomResponse{
		Room: room.Proto(),
	}
	return res, nil
}

func (s *classroomService) GetRoomsTotal(
	ctx context.Context, req *classroom.GetRoomsTotalRequest,
) (*classroom.GetRoomsTotalResponse, error) {
	if err := s.validator.GetRoomsTotal(req); err != nil {
		return nil, gRPCError(err)
	}

	total, err := s.db.Room.Count(ctx)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &classroom.GetRoomsTotalResponse{
		Total: total,
	}
	return res, nil
}

func (s *classroomService) UpdateRoomsTotal(
	ctx context.Context, req *classroom.UpdateRoomsTotalRequest,
) (*classroom.UpdateRoomsTotalResponse, error) {
	if err := s.validator.UpdateRoomsTotal(req); err != nil {
		return nil, gRPCError(err)
	}

	rooms := entity.NewRooms(int(req.Total))

	err := s.db.Room.Replace(ctx, rooms)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &classroom.UpdateRoomsTotalResponse{}
	return res, nil
}
