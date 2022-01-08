package request

import "github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"

type UpdateMySubjectRequest struct {
	SchoolType entity.SchoolType `json:"schoolType,omitempty"` // 校種
	SubjectIDs []int64           `json:"subjectIds,omitempty"` // 担当教科ID一覧
}

type UpdateMyMailRequest struct {
	Mail string `json:"mail,omitempty"` // メールアドレス
}

type UpdateMyPasswordRequest struct {
	Password             string `json:"password,omitempty"`             // パスワード
	PasswordConfirmation string `json:"passwordConfirmation,omitempty"` // パスワード(確認用)
}
