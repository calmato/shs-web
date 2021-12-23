package request

type UpdateMySubjectRequest struct {
	SchoolType int32   `json:"schoolType,omitempty"`
	SubjectIDs []int64 `json:"subjectIds,omitempty"`
}

type UpdateMyMailRequest struct {
	Mail string `json:"mail,omitempty"`
}

type UpdateMyPasswordRequest struct {
	Password             string `json:"password,omitempty"`
	PasswordConfirmation string `json:"passwordConfirmation,omitempty"`
}
