package jwt

import "time"

type JwtStructure interface {
	JwtHeader | JwtPayload
}

type JwtClaim map[string]string

type JwtHeader struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
}

func NewJwtHeader(algorithm string) *JwtHeader {
	return &JwtHeader{
		Algorithm: algorithm,
		Type:      "JWT",
	}
}

type JwtPayload struct {
	Expiration int64  `json:"exp"`
	Creation   int64  `json:"iat"`
	User       string `json:"user"`
	Role       string `json:"role"`
}

func NewJwtPayload(claim JwtClaim, d time.Duration) *JwtPayload {
	return &JwtPayload{
		Expiration: time.Now().Add(d).Unix(),
		Creation:   time.Now().Unix(),
		User:       claim["user"],
		Role:       claim["role"],
	}
}
