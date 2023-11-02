package jwt

import "time"

type JwtInterface interface {
	GenerateToken()
	RefreshToken()
}

type JwtToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	claim        JwtClaim
	key          *JwtKey
}

func NewJwtToken(user, role string) *JwtToken {
	return &JwtToken{
		claim: JwtClaim{
			"user": user,
			"role": role,
		},
		key: NewJwtKey(),
	}
}

func (t *JwtToken) Generate() {
	var sign *JwtSign
	// Generate JWT Access Token
	sign = NewJwtSign(t.claim, "HS256", time.Minute*30)
	sign.Generate(t.key.AccessKey)
	t.AccessToken = sign.Stringify()
	// Generate JWT Refresh Token
	sign = NewJwtSign(t.claim, "HS256", time.Minute*60)
	sign.Generate(t.key.RefreshKey)
	t.RefreshToken = sign.Stringify()
}
