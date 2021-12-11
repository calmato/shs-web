package entity

type Teacher struct {
	ID            string `json:"id"`            // 講師ID
	LastName      string `json:"lastName"`      // 姓
	FirstName     string `json:"firstName"`     // 名
	LastNameKana  string `json:"lastNameKana"`  // 姓(かな)
	FirstNameKana string `json:"firstNameKana"` // 名(かな)
	Mail          string `json:"mail"`          // メールアドレス
	Role          int32  `json:"role"`          // 権限
	CreatedAt     string `json:"createdAt"`     // 登録日時
	UpdatedAt     string `json:"updatedAt"`     // 更新日時
}

type Role int32

const (
	RoleUnknown       Role = 0
	RoleTeacher       Role = 1
	RoleAdministrator Role = 2
)
