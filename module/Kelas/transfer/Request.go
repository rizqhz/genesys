package transfer

type RequestParam struct {
	Kode string
}

type RequestQuery struct {
	Kode    string
	Nama    string
	Jurusan string
	Grade   string
	Tahun   int
}

type RequestBody struct {
	Kode    string `json:"kode" form:"kode"`
	Nama    string `json:"nama" form:"nama"`
	Jurusan string `json:"jurusan" form:"jurusan"`
	Grade   string `json:"grade" form:"grade"`
	Tahun   int    `json:"tahun" form:"tahun"`
}
