package handler

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/request"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/golang/mock/gomock"
)

func TestListSchedules(t *testing.T) {
	t.Parallel()
	schedules := []*classroom.Schedule{
		{
			Weekday:  int32(time.Sunday),
			IsClosed: true,
			Lessons:  []*classroom.Schedule_Lesson{},
		},
		{
			Weekday:  int32(time.Monday),
			IsClosed: false,
			Lessons: []*classroom.Schedule_Lesson{
				{StartTime: "1700", EndTime: "1830"},
				{StartTime: "1830", EndTime: "2000"},
			},
		},
	}
	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &classroom.ListSchedulesRequest{}
				out := &classroom.ListSchedulesResponse{Schedules: schedules}
				mocks.classroom.EXPECT().ListSchedules(gomock.Any(), in).Return(out, nil)
			},
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.SchedulesResponse{
					Schedules: entity.Schedules{
						{
							Weekday:  time.Sunday,
							IsClosed: true,
							Lessons:  entity.ScheduleLessons{},
						},
						{
							Weekday:  time.Monday,
							IsClosed: false,
							Lessons: entity.ScheduleLessons{
								{StartTime: "1700", EndTime: "1830"},
								{StartTime: "1830", EndTime: "2000"},
							},
						},
					},
				},
			},
		},
		{
			name: "failed to list schedules",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &classroom.ListSchedulesRequest{}
				mocks.classroom.EXPECT().ListSchedules(gomock.Any(), in).Return(nil, errmock)
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
			path := "/v1/schedules"
			req := newHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestUpdateSchedules(t *testing.T) {
	t.Parallel()
	schedules := []*classroom.ScheduleToUpdate{
		{
			Weekday:  int32(time.Sunday),
			IsClosed: false,
			Lessons: []*classroom.ScheduleToUpdate_Lesson{
				{StartTime: "1530", EndTime: "1700"},
				{StartTime: "1700", EndTime: "1830"},
				{StartTime: "1830", EndTime: "2000"},
				{StartTime: "2000", EndTime: "2130"},
			},
		},
		{
			Weekday:  int32(time.Monday),
			IsClosed: true,
			Lessons:  nil,
		},
		{
			Weekday:  int32(time.Tuesday),
			IsClosed: false,
			Lessons: []*classroom.ScheduleToUpdate_Lesson{
				{StartTime: "1700", EndTime: "1830"},
				{StartTime: "1830", EndTime: "2000"},
				{StartTime: "2000", EndTime: "2130"},
			},
		},
	}
	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		req    *request.UpdateSchedulesRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &classroom.UpdateSchedulesRequest{Schedules: schedules}
				out := &classroom.UpdateSchedulesResponse{}
				mocks.classroom.EXPECT().UpdateSchedules(gomock.Any(), in).Return(out, nil)
			},
			req: &request.UpdateSchedulesRequest{
				Schedules: []*request.ScheduleToUpdate{
					{
						Weekday:  time.Sunday,
						IsClosed: false,
						Lessons: []*request.ScheduleLesson{
							{StartTime: "1530", EndTime: "1700"},
							{StartTime: "1700", EndTime: "1830"},
							{StartTime: "1830", EndTime: "2000"},
							{StartTime: "2000", EndTime: "2130"},
						},
					},
					{
						Weekday:  time.Monday,
						IsClosed: true,
					},
					{
						Weekday:  time.Tuesday,
						IsClosed: false,
						Lessons: []*request.ScheduleLesson{
							{StartTime: "1700", EndTime: "1830"},
							{StartTime: "1830", EndTime: "2000"},
							{StartTime: "2000", EndTime: "2130"},
						},
					},
				},
			},
			expect: &testResponse{
				code: http.StatusNoContent,
			},
		},
		{
			name: "failed to update schedules",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &classroom.UpdateSchedulesRequest{Schedules: schedules}
				mocks.classroom.EXPECT().UpdateSchedules(gomock.Any(), in).Return(nil, errmock)
			},
			req: &request.UpdateSchedulesRequest{
				Schedules: []*request.ScheduleToUpdate{
					{
						Weekday:  time.Sunday,
						IsClosed: false,
						Lessons: []*request.ScheduleLesson{
							{StartTime: "1530", EndTime: "1700"},
							{StartTime: "1700", EndTime: "1830"},
							{StartTime: "1830", EndTime: "2000"},
							{StartTime: "2000", EndTime: "2130"},
						},
					},
					{
						Weekday:  time.Monday,
						IsClosed: true,
					},
					{
						Weekday:  time.Tuesday,
						IsClosed: false,
						Lessons: []*request.ScheduleLesson{
							{StartTime: "1700", EndTime: "1830"},
							{StartTime: "1830", EndTime: "2000"},
							{StartTime: "2000", EndTime: "2130"},
						},
					},
				},
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
			path := "/v1/schedules"
			req := newHTTPRequest(t, http.MethodPatch, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}
