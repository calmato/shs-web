package api

import (
	"context"
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/lesson/database"
	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/internal/lesson/validation"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func TestListLessons(t *testing.T) {
	t.Parallel()

	now := jst.Now()
	req := &lesson.ListLessonsRequest{
		ShiftSummaryId: 1,
		ShiftId:        1,
		TeacherId:      "teacherid",
		StudentId:      "studentid",
	}
	lessons := entity.Lessons{
		{
			ID:             1,
			ShiftSummaryID: 1,
			ShiftID:        1,
			SubjectID:      1,
			RoomID:         1,
			TeacherID:      "teacherid",
			StudentID:      "studentid",
			Notes:          "",
			CreatedAt:      now,
			UpdatedAt:      now,
		},
	}
	shifts := entity.Shifts{
		{
			ID:             1,
			ShiftSummaryID: 1,
			Date:           jst.Date(2022, 2, 1, 0, 0, 0, 0),
			StartTime:      "1700",
			EndTime:        "1830",
			CreatedAt:      now,
			UpdatedAt:      now,
		},
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *lesson.ListLessonsRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				params := &database.ListLessonsParams{
					ShiftSummaryID: 1,
					ShiftID:        1,
					TeacherID:      "teacherid",
					StudentID:      "studentid",
				}
				mocks.validator.EXPECT().ListLessons(req).Return(nil)
				mocks.db.Lesson.EXPECT().List(gomock.Any(), params).Return(lessons, nil)
				mocks.db.Lesson.EXPECT().Count(gomock.Any(), params).Return(int64(1), nil)
				mocks.db.Shift.EXPECT().MultiGet(ctx, []int64{1}).Return(shifts, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &lesson.ListLessonsResponse{
					Lessons: []*lesson.Lesson{
						{
							Id:             1,
							ShiftSummaryId: 1,
							ShiftId:        1,
							SubjectId:      1,
							RoomId:         1,
							TeacherId:      "teacherid",
							StudentId:      "studentid",
							Notes:          "",
							CreatedAt:      now.Unix(),
							UpdatedAt:      now.Unix(),
						},
					},
					Shifts: []*lesson.Shift{
						{
							Id:             1,
							ShiftSummaryId: 1,
							Date:           "20220201",
							StartTime:      "1700",
							EndTime:        "1830",
							CreatedAt:      now.Unix(),
							UpdatedAt:      now.Unix(),
						},
					},
					Total: 1,
				},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &lesson.ListLessonsRequest{}
				mocks.validator.EXPECT().ListLessons(req).Return(validation.ErrRequestValidation)
			},
			req: &lesson.ListLessonsRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to list lessons",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				params := &database.ListLessonsParams{
					ShiftSummaryID: 1,
					ShiftID:        1,
					TeacherID:      "teacherid",
					StudentID:      "studentid",
				}
				mocks.validator.EXPECT().ListLessons(req).Return(nil)
				mocks.db.Lesson.EXPECT().List(gomock.Any(), params).Return(nil, errmock)
				mocks.db.Lesson.EXPECT().Count(gomock.Any(), params).Return(int64(1), nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to count lessons",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				params := &database.ListLessonsParams{
					ShiftSummaryID: 1,
					ShiftID:        1,
					TeacherID:      "teacherid",
					StudentID:      "studentid",
				}
				mocks.validator.EXPECT().ListLessons(req).Return(nil)
				mocks.db.Lesson.EXPECT().List(gomock.Any(), params).Return(lessons, nil)
				mocks.db.Lesson.EXPECT().Count(gomock.Any(), params).Return(int64(0), errmock)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to count lessons",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				params := &database.ListLessonsParams{
					ShiftSummaryID: 1,
					ShiftID:        1,
					TeacherID:      "teacherid",
					StudentID:      "studentid",
				}
				mocks.validator.EXPECT().ListLessons(req).Return(nil)
				mocks.db.Lesson.EXPECT().List(gomock.Any(), params).Return(lessons, nil)
				mocks.db.Lesson.EXPECT().Count(gomock.Any(), params).Return(int64(1), nil)
				mocks.db.Shift.EXPECT().MultiGet(ctx, []int64{1}).Return(nil, errmock)
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
			return service.ListLessons(ctx, tt.req)
		}))
	}
}

func TestCreateLesson(t *testing.T) {
	t.Parallel()

	now := jst.Date(2022, 1, 15, 0, 0, 0, 0)
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
	shift := &entity.Shift{
		ID:             1,
		ShiftSummaryID: 1,
		Date:           jst.Date(2022, 2, 2, 0, 0, 0, 0),
		StartTime:      "1700",
		EndTime:        "1830",
		CreatedAt:      now,
		UpdatedAt:      now,
	}
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
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1)).Return(shift, nil)
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
					Shift: &lesson.Shift{
						Id:             1,
						ShiftSummaryId: 1,
						Date:           "20220202",
						StartTime:      "1700",
						EndTime:        "1830",
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
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
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1)).Return(shift, nil)
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
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1)).Return(shift, nil)
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
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1)).Return(shift, nil)
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
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1)).Return(shift, nil)
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
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1)).Return(nil, errmock)
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
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1)).Return(nil, errmock)
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
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1)).Return(shift, nil)
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

func TestUpdateLesson(t *testing.T) {
	t.Parallel()

	now := jst.Date(2022, 1, 15, 0, 0, 0, 0)
	req := &lesson.UpdateLessonRequest{
		LessonId:       1,
		ShiftSummaryId: 1,
		ShiftId:        1,
		SubjectId:      1,
		RoomId:         1,
		TeacherId:      "teacherid",
		StudentId:      "studentid",
		Notes:          "",
	}
	summary := &entity.ShiftSummary{ID: 1}
	shift := &entity.Shift{
		ID:             1,
		ShiftSummaryID: 1,
		Date:           jst.Date(2022, 2, 2, 0, 0, 0, 0),
		StartTime:      "1700",
		EndTime:        "1830",
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	l := &entity.Lesson{
		ID:             1,
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
		req    *lesson.UpdateLessonRequest
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
				mocks.validator.EXPECT().UpdateLesson(req).Return(nil)
				mocks.classroom.EXPECT().GetSubject(gomock.Any(), subjectIn).Return(subjectOut, nil)
				mocks.classroom.EXPECT().GetRoom(gomock.Any(), roomIn).Return(roomOut, nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1), "id").Return(summary, nil)
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1)).Return(shift, nil)
				mocks.db.Lesson.EXPECT().Update(ctx, int64(1), l).Return(nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &lesson.UpdateLessonResponse{},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &lesson.UpdateLessonRequest{}
				mocks.validator.EXPECT().UpdateLesson(req).Return(validation.ErrRequestValidation)
			},
			req: &lesson.UpdateLessonRequest{},
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
				mocks.validator.EXPECT().UpdateLesson(req).Return(nil)
				mocks.classroom.EXPECT().GetSubject(gomock.Any(), subjectIn).Return(nil, errmock)
				mocks.classroom.EXPECT().GetRoom(gomock.Any(), roomIn).Return(roomOut, nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1), "id").Return(summary, nil)
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1)).Return(shift, nil)
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
				mocks.validator.EXPECT().UpdateLesson(req).Return(nil)
				mocks.classroom.EXPECT().GetSubject(gomock.Any(), subjectIn).Return(subjectOut, nil)
				mocks.classroom.EXPECT().GetRoom(gomock.Any(), roomIn).Return(nil, errmock)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1), "id").Return(summary, nil)
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1)).Return(shift, nil)
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
				mocks.validator.EXPECT().UpdateLesson(req).Return(nil)
				mocks.classroom.EXPECT().GetSubject(gomock.Any(), subjectIn).Return(subjectOut, nil)
				mocks.classroom.EXPECT().GetRoom(gomock.Any(), roomIn).Return(roomOut, nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(nil, errmock)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1), "id").Return(summary, nil)
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1)).Return(shift, nil)
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
				mocks.validator.EXPECT().UpdateLesson(req).Return(nil)
				mocks.classroom.EXPECT().GetSubject(gomock.Any(), subjectIn).Return(subjectOut, nil)
				mocks.classroom.EXPECT().GetRoom(gomock.Any(), roomIn).Return(roomOut, nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(nil, errmock)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1), "id").Return(summary, nil)
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1)).Return(shift, nil)
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
				mocks.validator.EXPECT().UpdateLesson(req).Return(nil)
				mocks.classroom.EXPECT().GetSubject(gomock.Any(), subjectIn).Return(subjectOut, nil)
				mocks.classroom.EXPECT().GetRoom(gomock.Any(), roomIn).Return(roomOut, nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1), "id").Return(summary, nil)
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1)).Return(nil, errmock)
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
				mocks.validator.EXPECT().UpdateLesson(req).Return(nil)
				mocks.classroom.EXPECT().GetSubject(gomock.Any(), subjectIn).Return(subjectOut, nil)
				mocks.classroom.EXPECT().GetRoom(gomock.Any(), roomIn).Return(roomOut, nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1), "id").Return(summary, nil)
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1)).Return(nil, errmock)
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
				mocks.validator.EXPECT().UpdateLesson(req).Return(nil)
				mocks.classroom.EXPECT().GetSubject(gomock.Any(), subjectIn).Return(subjectOut, nil)
				mocks.classroom.EXPECT().GetRoom(gomock.Any(), roomIn).Return(roomOut, nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1), "id").Return(summary, nil)
				mocks.db.Shift.EXPECT().Get(gomock.Any(), int64(1)).Return(shift, nil)
				mocks.db.Lesson.EXPECT().Update(ctx, int64(1), l).Return(errmock)
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
			return service.UpdateLesson(ctx, tt.req)
		}))
	}
}

func TestDeleteLesson(t *testing.T) {
	t.Parallel()
	req := &lesson.DeleteLessonRequest{
		LessonId: 1,
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *lesson.DeleteLessonRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().DeleteLesson(req).Return(nil)
				mocks.db.Lesson.EXPECT().Delete(ctx, int64(1)).Return(nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &lesson.DeleteLessonRequest{}
				mocks.validator.EXPECT().DeleteLesson(req).Return(validation.ErrRequestValidation)
			},
			req: &lesson.DeleteLessonRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to delete lesson",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().DeleteLesson(req).Return(nil)
				mocks.db.Lesson.EXPECT().Delete(ctx, int64(1)).Return(errmock)
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
			return service.DeleteLesson(ctx, tt.req)
		}))
	}
}
