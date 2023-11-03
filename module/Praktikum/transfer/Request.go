package transfer

type RequestParam struct {
	ID string
}

type RequestQuery struct {
	ID            string
	MataPraktikum string
	Ruangan       string
	Kelas         string
	Jadwal        int
}

type RequestBody struct {
	ID            string `json:"praktikum" form:"praktikum"`
	MataPraktikum string `json:"mata_praktikum" form:"mata_praktikum"`
	Ruangan       string `json:"ruangan" form:"ruangan"`
	Kelas         string `json:"kelas" form:"kelas"`
	Jadwal        int    `json:"jadwal" form:"jadwal"`
}
