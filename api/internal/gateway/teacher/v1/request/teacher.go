package request

type CreateTeacherRequest struct {
	LastName             string `json:"lastName,omitempty"`             // 姓
	FirstName            string `json:"firstName,omitempty"`            // 名
	LastNameKana         string `json:"lastNameKana,omitempty"`         // 姓(かな)
	FirstNameKana        string `json:"firstNameKana,omitempty"`        // 名(かな)
	Mail                 string `json:"mail,omitempty"`                 // メールアドレス
	Password             string `json:"password,omitempty"`             // パスワード
	PasswordConfirmation string `json:"passwordConfirmation,omitempty"` // パスワード(確認用)
	Role                 int32  `json:"role,omitempty"`                 // 権限
}

type UpdateTeacherMailRequest struct {
	Mail string `json:"mail,omitempty"` // メールアドレス
}

type UpdateTeacherPasswordRequest struct {
	Password             string `json:"password,omitempty"`             // パスワード
	PasswordConfirmation string `json:"passwordConfirmation,omitempty"` // パスワード(確認用)
}
