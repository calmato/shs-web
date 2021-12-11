package request

type CreateTeacherRequest struct {
	LastName             string `json:"lastName,omitempty"`
	FirstName            string `json:"firstName,omitempty"`
	LastNameKana         string `json:"lastNameKana,omitempty"`
	FirstNameKana        string `json:"firstNameKana,omitempty"`
	Mail                 string `json:"mail,omitempty"`
	Password             string `json:"password,omitempty"`
	PasswordConfirmation string `json:"passwordConfirmation,omitempty"`
	Role                 int32  `json:"role,omitempty"`
}
