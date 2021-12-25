package request

type UpdateMySubjectRequest struct {
	SchoolType int32   `json:"schoolType,omitempty"` // 校種
	SubjectIDs []int64 `json:"subjectIds,omitempty"` // 担当教科ID一覧
}

type UpdateMyMailRequest struct {
	Mail string `json:"mail,omitempty"` // メールアドレス
}

type UpdateMyPasswordRequest struct {
	Password             string `json:"password,omitempty"`             // パスワード
	PasswordConfirmation string `json:"passwordConfirmation,omitempty"` // パスワード(確認用)
}
