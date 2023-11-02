package transfer

type Response struct {
	ID      int    `json:"jadwal_id"`
	Asisten string `json:"asisten"`
	Hari    string `json:"hari"`
	Jam     string `json:"jam"`
}
