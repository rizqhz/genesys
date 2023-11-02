package transfer

type RegisterResponse struct {
	UserID     int                `json:"user_id"`
	Nama       string             `json:"nama"`
	Email      string             `json:"email,omitempty"`
	Telepon    string             `json:"telepon,omitempty"`
	Credential CredentialResponse `json:"credential,omitempty"`
}

type CredentialResponse struct {
	Usercode string `json:"usercode,omitempty"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role,omitempty"`
	Token    string `json:"token,omitempty"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserResponse struct {
	ID      int    `json:"id"`
	Nama    string `json:"nama"`
	Alamat  string `json:"alamat"`
	Email   string `json:"email"`
	Telepon string `json:"telepon"`
	Foto    string `json:"foto"`
}
