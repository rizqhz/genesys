package transfer

type RequestParam struct {
	NIAS string
}

type RequestQuery struct {
	NIAS    string
	Nama    string
	Jabatan string
}

type RequestBody struct {
	NIAS    string `json:"nias" form:"nias"`
	Nama    string `json:"nama" form:"nama"`
	Jabatan string `json:"jabatan" form:"jabatan"`
}
