package transfer

type RegisterRequestBody struct {
	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email"`
	Telepon  string `json:"telepon" form:"telepon"`
	Usercode string `json:"usercode" form:"usercode"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
}

type LoginRequestBody struct {
	Usercode string `json:"usercode" form:"usercode"`
	Password string `json:"password" form:"password"`
}

type UserRequestBody struct {
	Nama    string `json:"nama" form:"nama"`
	Alamat  string `json:"alamat" form:"alamat"`
	Email   string `json:"email" form:"email"`
	Telepon string `json:"telepon" form:"telepon"`
	Foto    string
}
