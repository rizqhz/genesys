package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"time"

	"github.com/rizghz/genesys/internal/helpers"
)

type JwtSign struct {
	Header    string
	Payload   string
	Signature string
}

func NewJwtSign(claim JwtClaim, algorithm string, duration time.Duration) *JwtSign {
	header := NewJwtHeader(algorithm)
	payload := NewJwtPayload(claim, duration)
	return &JwtSign{
		Header:  helpers.JwtEncode[JwtHeader](header),
		Payload: helpers.JwtEncode[JwtPayload](payload),
	}
}

func (s *JwtSign) Generate(key string) {
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(s.Header + "." + s.Payload))
	result := hash.Sum(nil)
	encoder := base64.RawURLEncoding
	s.Signature = encoder.EncodeToString(result)
}

func (s *JwtSign) Stringify() string {
	return s.Header + "." + s.Payload + "." + s.Signature
}
