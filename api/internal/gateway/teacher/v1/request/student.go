package request

import "github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"

type CreateStudentRequest struct {
	LastName             string            `json:"lastName,omitempty"`             // 姓
	FirstName            string            `json:"firstName,omitempty"`            // 名
	LastNameKana         string            `json:"lastNameKana,omitempty"`         // 姓(かな)
	FirstNameKana        string            `json:"firstNameKana,omitempty"`        // 名(かな)
	Mail                 string            `json:"mail,omitempty"`                 // メールアドレス
	Password             string            `json:"password,omitempty"`             // パスワード
	PasswordConfirmation string            `json:"passwordConfirmation,omitempty"` // パスワード(確認用)
	SchoolType           entity.SchoolType `json:"schoolType,omitempty"`           // 校種
	Grade                int64             `json:"grade,omitempty"`                // 学年
}

type UpdateStudentMailRequest struct {
	Mail string `json:"mail,omitempty"` // メールアドレス
}

type UpdateStudentPasswordRequest struct {
	Password             string `json:"password,omitempty"`             // パスワード
	PasswordConfirmation string `json:"passwordConfirmation,omitempty"` // パスワード(確認用)
}

type UpdateStudentSubjectRequest struct {
	SchoolType entity.SchoolType `json:"schoolType,omitempty"` // 校種
	SubjectIDs []int64           `json:"subjectIds,omitempty"` // 受講教科ID一覧
}
