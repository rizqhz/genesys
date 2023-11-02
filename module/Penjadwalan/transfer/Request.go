package transfer

type RequestParam struct {
	ID int
}

type RequestQuery struct {
	Asisten string
	Hari    string
	Jam     string
}

type RequestBody struct {
	Asisten string `json:"asisten" form:"asisten"`
	Hari    string `json:"hari" form:"hari"`
	Jam     string `json:"jam" form:"jam"`
}
