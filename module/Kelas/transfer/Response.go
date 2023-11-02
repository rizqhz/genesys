package transfer

type Response struct {
	Kode    string `json:"kode"`
	Nama    string `json:"nama"`
	Jurusan string `json:"jurusan"`
	Grade   string `json:"grade"`
	Tahun   int    `json:"tahun"`
}
