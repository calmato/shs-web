package request

type UpdateMySubjectRequest struct {
	SchoolType int32   `json:"schoolType,omitempty"`
	SubjectIDs []int64 `json:"subjectIds,omitempty"`
}
