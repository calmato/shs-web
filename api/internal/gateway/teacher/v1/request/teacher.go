package request

import "github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"

type CreateTeacherRequest struct {
	LastName             string      `json:"lastName,omitempty"`             // 姓
	FirstName            string      `json:"firstName,omitempty"`            // 名
	LastNameKana         string      `json:"lastNameKana,omitempty"`         // 姓(かな)
	FirstNameKana        string      `json:"firstNameKana,omitempty"`        // 名(かな)
	Mail                 string      `json:"mail,omitempty"`                 // メールアドレス
	Password             string      `json:"password,omitempty"`             // パスワード
	PasswordConfirmation string      `json:"passwordConfirmation,omitempty"` // パスワード(確認用)
	Role                 entity.Role `json:"role,omitempty"`                 // 権限
}

type UpdateTeacherMailRequest struct {
	Mail string `json:"mail,omitempty"` // メールアドレス
}

type UpdateTeacherPasswordRequest struct {
	Password             string `json:"password,omitempty"`             // パスワード
	PasswordConfirmation string `json:"passwordConfirmation,omitempty"` // パスワード(確認用)
}

type UpdateTeacherRoleRequest struct {
	Role entity.Role `json:"role,omitempty"` // 権限
}

type UpdateTeacherSubjectRequest struct {
	SchoolType entity.SchoolType `json:"schoolType,omitempty"` // 校種
	SubjectIDs []int64           `json:"subjectIds,omitempty"` // 担当教科ID一覧
}
