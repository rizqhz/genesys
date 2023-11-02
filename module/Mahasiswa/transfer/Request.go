package transfer

type RequestParam struct {
	NPM string
}

type RequestQuery struct {
	NPm   string
	Nama  string
	Kelas string
}

type RequestBody struct {
	NPM   string `json:"npm" form:"npm"`
	Nama  string `json:"nama" form:"nama"`
	Kelas string `json:"kelas" form:"kelas"`
}
