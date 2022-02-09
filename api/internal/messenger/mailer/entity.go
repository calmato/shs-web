package mailer

import "fmt"

// 動的テンプレートID
const EmailIDLessonDecided = "lesson-decided" // 授業スケジュール確定

// Personalization - 送信先情報
type Personalization struct {
	Name          string                 // 送信先名称
	Address       string                 // 送信先メールアドレス
	Type          AddressType            // 宛先タイプ
	Substitutions map[string]interface{} // 動的コンテンツ
}

// AddressType - メール宛先種別
type AddressType int32

// Content - メッセージ内容
type Content struct {
	ContentType string
	Value       string
}

const (
	AddressTypeTo AddressType = iota
	AddressTypeCC
	AddressTypeBCC
)

type SendGridError struct {
	Code    int
	Message string `json:"message"`
	Field   string `json:"field"`
	Help    string `json:"help"`
}

func (e SendGridError) Error() string {
	return fmt.Sprintf("mail: failed to SendGrid. code=%d, %s=%s", e.Code, e.Field, e.Message)
}

/**
 * ---------------------------
 * Other
 * ---------------------------
 */

// Substitutions - メール動的コンテンツの生成
func NewSubstitutions(params map[string]string) map[string]interface{} {
	res := make(map[string]interface{}, len(params))
	for k, v := range params {
		res[k] = v
	}
	return res
}
