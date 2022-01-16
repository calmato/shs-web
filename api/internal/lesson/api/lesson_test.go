package api

import (
	"context"
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/internal/lesson/validation"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func TestCreateLesson(t *testing.T) {
	t.Parallel()

	req := &lesson.CreateLessonRequest{
		ShiftSummaryId: 1,
		ShiftId:        1,
		SubjectId:      1,
		RoomId:         1,
		TeacherId:      "teacherid",
		StudentId:      "studentid",
		Notes:          "",
	}
	summary := &entity.ShiftSummary{ID: 1}
	shift := &entity.Shift{ID: 1}
	l := &entity.Lesson{
		ShiftSummaryID: 1,
		ShiftID:        1,
		SubjectID:      1,
		RoomID:         1,
		TeacherID:      "teacherid",
		StudentID:      "studentid",
		Notes:          "",
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *lesson.CreateLessonRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				subjectIn := &classroom.GetSubjectRequest{Id: 1}
				subjectOut := &classroom.GetSubjectResponse{Subject: &classroom.Subject{Id: 1}}
				roomIn := &classroom.GetRoomRequest{Id: 1}
				roomOut := &classroom.GetRoomResponse{Room: &classroom.Room{Id: 1}}
				teacherIn := &user.GetTeacherRequest{Id: "teacherid"}
				teacherOut := &user.GetTeacherResponse{Teacher: &user.Teacher{Id: "teacherid"}}
				studentIn := &user.GetStudentRequest{Id: "studentid"}
				studentOut := &user.GetStudentResponse{Student: &user.Student{Id: "studentid"}}
				mocks.validator.EXPECT().CreateLesson(req).Return(nil)
				mocks.classroom.EXPECT().GetSubject(gomock.Any(), subjectIn).Return(subjectOut, nil)
				mocks.classroom.EXPECT().GetRoom(gomock.Any(), roomIn).Return(roomOut, nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1), "id").Return(summary, nil)
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1), "id").Return(shift, nil)
				mocks.db.Lesson.EXPECT().Create(ctx, l).Return(nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &lesson.CreateLessonResponse{
					Lesson: &lesson.Lesson{
						Id:             0,
						ShiftSummaryId: 1,
						ShiftId:        1,
						SubjectId:      1,
						RoomId:         1,
						TeacherId:      "teacherid",
						StudentId:      "studentid",
						Notes:          "",
						CreatedAt:      time.Time{}.Unix(),
						UpdatedAt:      time.Time{}.Unix(),
					},
				},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &lesson.CreateLessonRequest{}
				mocks.validator.EXPECT().CreateLesson(req).Return(validation.ErrRequestValidation)
			},
			req: &lesson.CreateLessonRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to get subject",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				subjectIn := &classroom.GetSubjectRequest{Id: 1}
				roomIn := &classroom.GetRoomRequest{Id: 1}
				roomOut := &classroom.GetRoomResponse{Room: &classroom.Room{Id: 1}}
				teacherIn := &user.GetTeacherRequest{Id: "teacherid"}
				teacherOut := &user.GetTeacherResponse{Teacher: &user.Teacher{Id: "teacherid"}}
				studentIn := &user.GetStudentRequest{Id: "studentid"}
				studentOut := &user.GetStudentResponse{Student: &user.Student{Id: "studentid"}}
				mocks.validator.EXPECT().CreateLesson(req).Return(nil)
				mocks.classroom.EXPECT().GetSubject(gomock.Any(), subjectIn).Return(nil, errmock)
				mocks.classroom.EXPECT().GetRoom(gomock.Any(), roomIn).Return(roomOut, nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1), "id").Return(summary, nil)
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1), "id").Return(shift, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to get room",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				subjectIn := &classroom.GetSubjectRequest{Id: 1}
				subjectOut := &classroom.GetSubjectResponse{Subject: &classroom.Subject{Id: 1}}
				roomIn := &classroom.GetRoomRequest{Id: 1}
				teacherIn := &user.GetTeacherRequest{Id: "teacherid"}
				teacherOut := &user.GetTeacherResponse{Teacher: &user.Teacher{Id: "teacherid"}}
				studentIn := &user.GetStudentRequest{Id: "studentid"}
				studentOut := &user.GetStudentResponse{Student: &user.Student{Id: "studentid"}}
				mocks.validator.EXPECT().CreateLesson(req).Return(nil)
				mocks.classroom.EXPECT().GetSubject(gomock.Any(), subjectIn).Return(subjectOut, nil)
				mocks.classroom.EXPECT().GetRoom(gomock.Any(), roomIn).Return(nil, errmock)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1), "id").Return(summary, nil)
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1), "id").Return(shift, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to get teacher",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				subjectIn := &classroom.GetSubjectRequest{Id: 1}
				subjectOut := &classroom.GetSubjectResponse{Subject: &classroom.Subject{Id: 1}}
				roomIn := &classroom.GetRoomRequest{Id: 1}
				roomOut := &classroom.GetRoomResponse{Room: &classroom.Room{Id: 1}}
				teacherIn := &user.GetTeacherRequest{Id: "teacherid"}
				studentIn := &user.GetStudentRequest{Id: "studentid"}
				studentOut := &user.GetStudentResponse{Student: &user.Student{Id: "studentid"}}
				mocks.validator.EXPECT().CreateLesson(req).Return(nil)
				mocks.classroom.EXPECT().GetSubject(gomock.Any(), subjectIn).Return(subjectOut, nil)
				mocks.classroom.EXPECT().GetRoom(gomock.Any(), roomIn).Return(roomOut, nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(nil, errmock)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1), "id").Return(summary, nil)
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1), "id").Return(shift, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to get student",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				subjectIn := &classroom.GetSubjectRequest{Id: 1}
				subjectOut := &classroom.GetSubjectResponse{Subject: &classroom.Subject{Id: 1}}
				roomIn := &classroom.GetRoomRequest{Id: 1}
				roomOut := &classroom.GetRoomResponse{Room: &classroom.Room{Id: 1}}
				teacherIn := &user.GetTeacherRequest{Id: "teacherid"}
				teacherOut := &user.GetTeacherResponse{Teacher: &user.Teacher{Id: "teacherid"}}
				studentIn := &user.GetStudentRequest{Id: "studentid"}
				mocks.validator.EXPECT().CreateLesson(req).Return(nil)
				mocks.classroom.EXPECT().GetSubject(gomock.Any(), subjectIn).Return(subjectOut, nil)
				mocks.classroom.EXPECT().GetRoom(gomock.Any(), roomIn).Return(roomOut, nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(nil, errmock)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1), "id").Return(summary, nil)
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1), "id").Return(shift, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to get shift summary",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				subjectIn := &classroom.GetSubjectRequest{Id: 1}
				subjectOut := &classroom.GetSubjectResponse{Subject: &classroom.Subject{Id: 1}}
				roomIn := &classroom.GetRoomRequest{Id: 1}
				roomOut := &classroom.GetRoomResponse{Room: &classroom.Room{Id: 1}}
				teacherIn := &user.GetTeacherRequest{Id: "teacherid"}
				teacherOut := &user.GetTeacherResponse{Teacher: &user.Teacher{Id: "teacherid"}}
				studentIn := &user.GetStudentRequest{Id: "studentid"}
				studentOut := &user.GetStudentResponse{Student: &user.Student{Id: "studentid"}}
				mocks.validator.EXPECT().CreateLesson(req).Return(nil)
				mocks.classroom.EXPECT().GetSubject(gomock.Any(), subjectIn).Return(subjectOut, nil)
				mocks.classroom.EXPECT().GetRoom(gomock.Any(), roomIn).Return(roomOut, nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1), "id").Return(summary, nil)
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1), "id").Return(nil, errmock)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to get shift",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				subjectIn := &classroom.GetSubjectRequest{Id: 1}
				subjectOut := &classroom.GetSubjectResponse{Subject: &classroom.Subject{Id: 1}}
				roomIn := &classroom.GetRoomRequest{Id: 1}
				roomOut := &classroom.GetRoomResponse{Room: &classroom.Room{Id: 1}}
				teacherIn := &user.GetTeacherRequest{Id: "teacherid"}
				teacherOut := &user.GetTeacherResponse{Teacher: &user.Teacher{Id: "teacherid"}}
				studentIn := &user.GetStudentRequest{Id: "studentid"}
				studentOut := &user.GetStudentResponse{Student: &user.Student{Id: "studentid"}}
				mocks.validator.EXPECT().CreateLesson(req).Return(nil)
				mocks.classroom.EXPECT().GetSubject(gomock.Any(), subjectIn).Return(subjectOut, nil)
				mocks.classroom.EXPECT().GetRoom(gomock.Any(), roomIn).Return(roomOut, nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1), "id").Return(summary, nil)
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1), "id").Return(nil, errmock)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to create lesson",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				subjectIn := &classroom.GetSubjectRequest{Id: 1}
				subjectOut := &classroom.GetSubjectResponse{Subject: &classroom.Subject{Id: 1}}
				roomIn := &classroom.GetRoomRequest{Id: 1}
				roomOut := &classroom.GetRoomResponse{Room: &classroom.Room{Id: 1}}
				teacherIn := &user.GetTeacherRequest{Id: "teacherid"}
				teacherOut := &user.GetTeacherResponse{Teacher: &user.Teacher{Id: "teacherid"}}
				studentIn := &user.GetStudentRequest{Id: "studentid"}
				studentOut := &user.GetStudentResponse{Student: &user.Student{Id: "studentid"}}
				mocks.validator.EXPECT().CreateLesson(req).Return(nil)
				mocks.classroom.EXPECT().GetSubject(gomock.Any(), subjectIn).Return(subjectOut, nil)
				mocks.classroom.EXPECT().GetRoom(gomock.Any(), roomIn).Return(roomOut, nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1), "id").Return(summary, nil)
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1), "id").Return(shift, nil)
				mocks.db.Lesson.EXPECT().Create(ctx, l).Return(errmock)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, testGRPC(tt.setup, tt.expect, func(ctx context.Context, service *lessonService) (proto.Message, error) {
			return service.CreateLesson(ctx, tt.req)
		}))
	}
}
