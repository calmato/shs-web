package request

type UpdateMyMailRequest struct {
	Mail string `json:"mail,omitempty"` // メールアドレス
}

type UpdateMyPasswordRequest struct {
	Password             string `json:"password,omitempty"`             // パスワード
	PasswordConfirmation string `json:"passwordConfirmation,omitempty"` // パスワード(確認用)
}
