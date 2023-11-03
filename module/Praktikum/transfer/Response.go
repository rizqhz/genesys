package transfer

type Response struct {
	ID                string `json:"praktikum"`
	KodeMataPraktikum string `json:"mata_praktikum"`
	KodeRuangan       string `json:"ruangan"`
	KodeKelas         string `json:"kelas"`
	JadwalID          int    `json:"jadwal"`
}
