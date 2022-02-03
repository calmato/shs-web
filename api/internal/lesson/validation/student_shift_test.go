package validation

import (
	"testing"
	"time"

	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestListStudentSubmissionsByShiftSummaryIDs(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *lesson.ListStudentSubmissionsByShiftSummaryIDsRequest
		isErr bool
	}{
		{
			name: "success",
			req: &lesson.ListStudentSubmissionsByShiftSummaryIDsRequest{
				StudentId:       "studentid",
				ShiftSummaryIds: []int64{1, 2},
			},
			isErr: false,
		},
		{
			name: "StudentId is min_len",
			req: &lesson.ListStudentSubmissionsByShiftSummaryIDsRequest{
				StudentId:       "",
				ShiftSummaryIds: []int64{1, 2},
			},
			isErr: true,
		},
		{
			name: "ShiftSummaryIds is unique",
			req: &lesson.ListStudentSubmissionsByShiftSummaryIDsRequest{
				StudentId:       "studentid",
				ShiftSummaryIds: []int64{1, 1},
			},
			isErr: true,
		},
		{
			name: "ShiftSummaryIds is items.gt",
			req: &lesson.ListStudentSubmissionsByShiftSummaryIDsRequest{
				StudentId:       "studentid",
				ShiftSummaryIds: []int64{0},
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.ListStudentSubmissionsByShiftSummaryIDs(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestListStudentSubmissionsByStudentIDs(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *lesson.ListStudentSubmissionsByStudentIDsRequest
		isErr bool
	}{
		{
			name: "success",
			req: &lesson.ListStudentSubmissionsByStudentIDsRequest{
				StudentIds:     []string{"studentid"},
				ShiftSummaryId: 1,
			},
			isErr: false,
		},
		{
			name: "StudentIds is unique",
			req: &lesson.ListStudentSubmissionsByStudentIDsRequest{
				StudentIds:     []string{"studentid", "studentid"},
				ShiftSummaryId: 1,
			},
			isErr: true,
		},
		{
			name: "StudentIds is items.min_len",
			req: &lesson.ListStudentSubmissionsByStudentIDsRequest{
				StudentIds:     []string{""},
				ShiftSummaryId: 1,
			},
			isErr: true,
		},
		{
			name: "ShiftSummaryId is gt",
			req: &lesson.ListStudentSubmissionsByStudentIDsRequest{
				StudentIds:     []string{"studentid"},
				ShiftSummaryId: 0,
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.ListStudentSubmissionsByStudentIDs(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestListStudentShifts(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *lesson.ListStudentShiftsRequest
		isErr bool
	}{
		{
			name: "success",
			req: &lesson.ListStudentShiftsRequest{
				StudentIds:     []string{"studentid"},
				ShiftSummaryId: 1,
			},
			isErr: false,
		},
		{
			name: "StudentIds is unique",
			req: &lesson.ListStudentShiftsRequest{
				StudentIds:     []string{"studentid", "studentid"},
				ShiftSummaryId: 1,
			},
			isErr: true,
		},
		{
			name: "StudentIds is items.min_len",
			req: &lesson.ListStudentShiftsRequest{
				StudentIds:     []string{""},
				ShiftSummaryId: 1,
			},
			isErr: true,
		},
		{
			name: "ShiftSummaryId is gt",
			req: &lesson.ListStudentShiftsRequest{
				StudentIds:     []string{"studentid"},
				ShiftSummaryId: 0,
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.ListStudentShifts(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestGetStudentShifts(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *lesson.GetStudentShiftsRequest
		isErr bool
	}{
		{
			name: "success",
			req: &lesson.GetStudentShiftsRequest{
				StudentId:      "studentid",
				ShiftSummaryId: 1,
			},
			isErr: false,
		},
		{
			name: "StudentId is min_len",
			req: &lesson.GetStudentShiftsRequest{
				StudentId:      "",
				ShiftSummaryId: 1,
			},
			isErr: true,
		},
		{
			name: "ShiftSummaryId is gt",
			req: &lesson.GetStudentShiftsRequest{
				StudentId:      "studentid",
				ShiftSummaryId: 0,
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.GetStudentShifts(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestUpsertStudentShifts(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *lesson.UpsertStudentShiftsRequest
		isErr bool
	}{
		{
			name: "success",
			req: &lesson.UpsertStudentShiftsRequest{
				StudentId:      "studentid",
				ShiftSummaryId: 1,
				ShiftIds:       []int64{1, 2},
				Decided:        true,
				Lessons:        []*lesson.StudentSuggestedLesson{{SubjectId: 1, Total: 4}},
			},
			isErr: false,
		},
		{
			name: "StudentId is min_len",
			req: &lesson.UpsertStudentShiftsRequest{
				StudentId:      "",
				ShiftSummaryId: 1,
				ShiftIds:       []int64{1, 2},
				Decided:        true,
				Lessons:        []*lesson.StudentSuggestedLesson{{SubjectId: 1, Total: 4}},
			},
			isErr: true,
		},
		{
			name: "ShiftSummaryId is gt",
			req: &lesson.UpsertStudentShiftsRequest{
				StudentId:      "studentid",
				ShiftSummaryId: 0,
				ShiftIds:       []int64{1, 2},
				Decided:        true,
				Lessons:        []*lesson.StudentSuggestedLesson{{SubjectId: 1, Total: 4}},
			},
			isErr: true,
		},
		{
			name: "ShiftIds is unique",
			req: &lesson.UpsertStudentShiftsRequest{
				StudentId:      "studentid",
				ShiftSummaryId: 1,
				ShiftIds:       []int64{1, 1},
				Decided:        true,
				Lessons:        []*lesson.StudentSuggestedLesson{{SubjectId: 1, Total: 4}},
			},
			isErr: true,
		},
		{
			name: "ShiftIds is items.gt",
			req: &lesson.UpsertStudentShiftsRequest{
				StudentId:      "studentid",
				ShiftSummaryId: 1,
				ShiftIds:       []int64{0},
				Decided:        true,
				Lessons:        []*lesson.StudentSuggestedLesson{{SubjectId: 1, Total: 4}},
			},
			isErr: true,
		},
		{
			name: "Lessons.SubjectId is gt",
			req: &lesson.UpsertStudentShiftsRequest{
				StudentId:      "studentid",
				ShiftSummaryId: 1,
				ShiftIds:       []int64{0},
				Decided:        true,
				Lessons:        []*lesson.StudentSuggestedLesson{{SubjectId: 0, Total: 4}},
			},
			isErr: true,
		},
		{
			name: "Lessons.Total is gte",
			req: &lesson.UpsertStudentShiftsRequest{
				StudentId:      "studentid",
				ShiftSummaryId: 1,
				ShiftIds:       []int64{0},
				Decided:        true,
				Lessons:        []*lesson.StudentSuggestedLesson{{SubjectId: 1, Total: -1}},
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.UpsertStudentShifts(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestGetStudentShiftTemplate(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *lesson.GetStudentShiftTemplateRequest
		isErr bool
	}{
		{
			name: "success",
			req: &lesson.GetStudentShiftTemplateRequest{
				StudentId: "studentid",
			},
			isErr: false,
		},
		{
			name: "StudentId is min_len",
			req: &lesson.GetStudentShiftTemplateRequest{
				StudentId: "",
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.GetStudentShiftTemplate(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestUpsertStudentShiftTemplate(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *lesson.UpsertStudentShiftTemplateRequest
		isErr bool
	}{
		{
			name: "success",
			req: &lesson.UpsertStudentShiftTemplateRequest{
				StudentId: "studentid",
				Template: &lesson.StudentShiftTemplateToUpsert{
					Schedules: []*lesson.StudentShiftTemplateToUpsert_Schedule{
						{
							Weekday: int32(time.Sunday),
							Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{{StartTime: "1700", EndTime: "1830"}},
						},
					},
					SuggestedLessons: []*lesson.StudentSuggestedLesson{
						{SubjectId: 1, Total: 4},
					},
				},
			},
			isErr: false,
		},
		{
			name: "StudentId is min_len",
			req: &lesson.UpsertStudentShiftTemplateRequest{
				StudentId: "",
				Template: &lesson.StudentShiftTemplateToUpsert{
					Schedules: []*lesson.StudentShiftTemplateToUpsert_Schedule{
						{
							Weekday: int32(time.Sunday),
							Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{{StartTime: "1700", EndTime: "1830"}},
						},
					},
					SuggestedLessons: []*lesson.StudentSuggestedLesson{
						{SubjectId: 1, Total: 4},
					},
				},
			},
			isErr: true,
		},
		{
			name: "Template is required",
			req: &lesson.UpsertStudentShiftTemplateRequest{
				StudentId: "studentid",
				Template:  nil,
			},
			isErr: true,
		},
		{
			name: "Template.Schedules is max_items",
			req: &lesson.UpsertStudentShiftTemplateRequest{
				StudentId: "studentid",
				Template: &lesson.StudentShiftTemplateToUpsert{
					Schedules: []*lesson.StudentShiftTemplateToUpsert_Schedule{
						{Weekday: int32(time.Sunday), Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{}},
						{Weekday: int32(time.Monday), Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{}},
						{Weekday: int32(time.Tuesday), Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{}},
						{Weekday: int32(time.Wednesday), Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{}},
						{Weekday: int32(time.Thursday), Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{}},
						{Weekday: int32(time.Friday), Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{}},
						{Weekday: int32(time.Saturday), Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{}},
						{Weekday: int32(time.Sunday), Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{}},
					},
					SuggestedLessons: []*lesson.StudentSuggestedLesson{
						{SubjectId: 1, Total: 4},
					},
				},
			},
			isErr: true,
		},
		{
			name: "Template.Schedules.Weekday is gte",
			req: &lesson.UpsertStudentShiftTemplateRequest{
				StudentId: "studentid",
				Template: &lesson.StudentShiftTemplateToUpsert{
					Schedules: []*lesson.StudentShiftTemplateToUpsert_Schedule{
						{Weekday: -1, Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{}},
					},
					SuggestedLessons: []*lesson.StudentSuggestedLesson{
						{SubjectId: 1, Total: 4},
					},
				},
			},
			isErr: true,
		},
		{
			name: "Template.Schedules.Weekday is lte",
			req: &lesson.UpsertStudentShiftTemplateRequest{
				StudentId: "studentid",
				Template: &lesson.StudentShiftTemplateToUpsert{
					Schedules: []*lesson.StudentShiftTemplateToUpsert_Schedule{
						{Weekday: 7, Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{}},
					},
					SuggestedLessons: []*lesson.StudentSuggestedLesson{
						{SubjectId: 1, Total: 4},
					},
				},
			},
			isErr: true,
		},
		{
			name: "Template.Schedules.Lessons.StartTime is len",
			req: &lesson.UpsertStudentShiftTemplateRequest{
				StudentId: "studentid",
				Template: &lesson.StudentShiftTemplateToUpsert{
					Schedules: []*lesson.StudentShiftTemplateToUpsert_Schedule{
						{
							Weekday: int32(time.Sunday),
							Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{{StartTime: "", EndTime: "1830"}},
						},
					},
					SuggestedLessons: []*lesson.StudentSuggestedLesson{
						{SubjectId: 1, Total: 4},
					},
				},
			},
			isErr: true,
		},
		{
			name: "Template.Schedules.Lessons.StartTime is pattern",
			req: &lesson.UpsertStudentShiftTemplateRequest{
				StudentId: "studentid",
				Template: &lesson.StudentShiftTemplateToUpsert{
					Schedules: []*lesson.StudentShiftTemplateToUpsert_Schedule{
						{
							Weekday: int32(time.Sunday),
							Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{{StartTime: "aaaa", EndTime: "1830"}},
						},
					},
					SuggestedLessons: []*lesson.StudentSuggestedLesson{
						{SubjectId: 1, Total: 4},
					},
				},
			},
			isErr: true,
		},
		{
			name: "Template.Schedules.Lessons.EndTime is len",
			req: &lesson.UpsertStudentShiftTemplateRequest{
				StudentId: "studentid",
				Template: &lesson.StudentShiftTemplateToUpsert{
					Schedules: []*lesson.StudentShiftTemplateToUpsert_Schedule{
						{
							Weekday: int32(time.Sunday),
							Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{{StartTime: "1700", EndTime: ""}},
						},
					},
					SuggestedLessons: []*lesson.StudentSuggestedLesson{
						{SubjectId: 1, Total: 4},
					},
				},
			},
			isErr: true,
		},
		{
			name: "Template.Schedules.Lessons.StartTime is pattern",
			req: &lesson.UpsertStudentShiftTemplateRequest{
				StudentId: "studentid",
				Template: &lesson.StudentShiftTemplateToUpsert{
					Schedules: []*lesson.StudentShiftTemplateToUpsert_Schedule{
						{
							Weekday: int32(time.Sunday),
							Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{{StartTime: "1700", EndTime: "aaaa"}},
						},
					},
					SuggestedLessons: []*lesson.StudentSuggestedLesson{
						{SubjectId: 1, Total: 4},
					},
				},
			},
			isErr: true,
		},
		{
			name: "Template.SuggestedLessons.SubjectId is gt",
			req: &lesson.UpsertStudentShiftTemplateRequest{
				StudentId: "studentid",
				Template: &lesson.StudentShiftTemplateToUpsert{
					Schedules: []*lesson.StudentShiftTemplateToUpsert_Schedule{
						{
							Weekday: int32(time.Sunday),
							Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{{StartTime: "1700", EndTime: "1830"}},
						},
					},
					SuggestedLessons: []*lesson.StudentSuggestedLesson{
						{SubjectId: 0, Total: 4},
					},
				},
			},
			isErr: true,
		},
		{
			name: "Template.SuggestedLessons.Total is gte",
			req: &lesson.UpsertStudentShiftTemplateRequest{
				StudentId: "studentid",
				Template: &lesson.StudentShiftTemplateToUpsert{
					Schedules: []*lesson.StudentShiftTemplateToUpsert_Schedule{
						{
							Weekday: int32(time.Sunday),
							Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{{StartTime: "1700", EndTime: "1830"}},
						},
					},
					SuggestedLessons: []*lesson.StudentSuggestedLesson{
						{SubjectId: 1, Total: -1},
					},
				},
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.UpsertStudentShiftTemplate(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}
