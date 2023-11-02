package transfer

type RequestParam struct {
	Kode string
}

type RequestQuery struct {
	Kode string
	Nama string
}

type RequestBody struct {
	Kode string `json:"kode" form:"kode"`
	Nama string `json:"nama" form:"nama"`
}
