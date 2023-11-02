package auth

type UserEntity struct {
	ID      int
	Nama    string
	Alamat  string
	Email   string
	Telepon string
	Foto    string
}

type CredentialEntity struct {
	ID       int
	Usercode string
	Password string
	Role     string
	Token    string
}
