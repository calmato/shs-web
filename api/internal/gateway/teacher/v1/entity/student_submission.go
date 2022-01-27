package entity

import (
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
)

type StudentSubmission struct {
	ShiftSummaryID   int64                   `json:"id"`               // 授業希望募集ID
	Year             int32                   `json:"year"`             // 年
	Month            int32                   `json:"month"`            // 月
	ShiftStatus      ShiftStatus             `json:"shiftStatus"`      // 授業希望募集ステータス
	SubmissionStatus StudentSubmissionStatus `json:"submissionStatus"` // 授業希望提出ステータス
	OpenAt           time.Time               `json:"openAt"`           // 募集開始日時
	EndAt            time.Time               `json:"endAt"`            // 募集締切日時
	CreatedAt        time.Time               `json:"createdAt"`        // 登録日時
	UpdatedAt        time.Time               `json:"updatedAt"`        // 更新日時
}

type StudentSubmissions []*StudentSubmission

type StudentSuggestedLesson struct {
	SubjectID int64 `json:"subjectId"` // 授業ID
	Total     int64 `json:"total"`     // 希望授業数
}

type StudentSuggestedLessons []*StudentSuggestedLesson

type StudentSubmissionDetail struct {
	Student               *Student `json:"student"`               // 生徒情報
	LessonTotal           int64    `json:"lessonTotal"`           // 登録受講授業数
	SuggestedLessonsTotal int64    `json:"suggestedLessonsTotal"` // 希望受講授業数
}

type StudentSubmissionDetails []*StudentSubmissionDetail

type StudentSubmissionStatus int32

const (
	StudentSubmissionStatusUnknown   StudentSubmissionStatus = 0
	StudentSubmissionStatusWaiting   StudentSubmissionStatus = 1
	StudentSubmissionStatusSubmitted StudentSubmissionStatus = 2
)

func NewStudentSubmission(summary *entity.ShiftSummary, submission *entity.StudentSubmission) *StudentSubmission {
	return &StudentSubmission{
		ShiftSummaryID:   summary.Id,
		Year:             summary.YearMonth / 100,
		Month:            summary.YearMonth % 100,
		ShiftStatus:      NewShiftStatus(summary.Status),
		SubmissionStatus: NewStudentSubmissionStatus(summary, submission),
		OpenAt:           jst.ParseFromUnix(summary.OpenAt),
		EndAt:            jst.ParseFromUnix(summary.EndAt),
		CreatedAt:        jst.ParseFromUnix(summary.CreatedAt),
		UpdatedAt:        jst.ParseFromUnix(summary.UpdatedAt),
	}
}

func NewStudentSubmissions(
	summaries entity.ShiftSummaries, submissions map[int64]*entity.StudentSubmission,
) StudentSubmissions {
	ss := make(StudentSubmissions, len(summaries))
	for i, s := range summaries {
		submission := submissions[s.Id] // null: 出勤不可
		ss[i] = NewStudentSubmission(s, submission)
	}
	return ss
}

func NewStudentSuggestedLesson(suggestedLesson *lesson.SuggestedLesson) *StudentSuggestedLesson {
	return &StudentSuggestedLesson{
		SubjectID: suggestedLesson.SubjectId,
		Total:     suggestedLesson.Total,
	}
}

func NewStudentSuggestedLessons(submission *entity.StudentSubmission) StudentSuggestedLessons {
	ls := make(StudentSuggestedLessons, len(submission.SuggestedLessons))
	for i := range submission.SuggestedLessons {
		ls[i] = NewStudentSuggestedLesson(submission.SuggestedLessons[i])
	}
	return ls
}

func NewStudentSubmissionDetail(
	student *entity.Student,
	subjects entity.Subjects,
	submission *entity.StudentSubmission,
	lessons entity.Lessons,
) *StudentSubmissionDetail {
	var suggestedLessonsTotal int64
	if submission != nil {
		for _, l := range submission.SuggestedLessons {
			suggestedLessonsTotal += l.Total
		}
	}
	return &StudentSubmissionDetail{
		Student:               NewStudent(student, subjects),
		LessonTotal:           int64(len(lessons)),
		SuggestedLessonsTotal: suggestedLessonsTotal,
	}
}

func NewStudentSubmissionDetails(
	students entity.Students,
	subjectsMap map[string]entity.Subjects,
	submissionMap map[string]*entity.StudentSubmission,
	lessonsMap map[string]entity.Lessons,
) StudentSubmissionDetails {
	ss := make(StudentSubmissionDetails, len(students))
	for i, student := range students {
		subjects := subjectsMap[student.Id]
		submission := submissionMap[student.Id]
		lessons := lessonsMap[student.Id]
		ss[i] = NewStudentSubmissionDetail(student, subjects, submission, lessons)
	}
	return ss
}

func NewStudentSubmissionStatus(
	summary *entity.ShiftSummary, submission *entity.StudentSubmission,
) StudentSubmissionStatus {
	if summary == nil {
		return StudentSubmissionStatusUnknown
	}
	if submission != nil && submission.Decided {
		return StudentSubmissionStatusSubmitted
	}
	return StudentSubmissionStatusWaiting
}
