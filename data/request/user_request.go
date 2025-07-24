package request

type RegisterUser struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUser struct {
	Name        string `json:"name"`
	ProfilePict string `json:"profile_pict"`
}
