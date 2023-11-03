package helpers

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func JwtEncode[T any](structure *T) (str string) {
	encoder := base64.RawStdEncoding
	json, err := json.Marshal(structure)
	if err != nil {
		log.Fatal("jwt.encode:", err.Error())
	}
	str = encoder.EncodeToString(json)
	return str
}

func JwtDecode[T any](str *string) (structure *T) {
	decoder := base64.RawStdEncoding
	data, err := decoder.DecodeString(*str)
	if err != nil {
		log.Fatal("jwt.decode:", err.Error())
	}
	structure = new(T)
	if err := json.Unmarshal(data, structure); err != nil {
		log.Fatal("jwt.decode:", err.Error())
	}
	return structure
}

func JwtValidate(token []string, key string) bool {
	buffer := strings.Join(token, ".")
	res, err := jwt.Parse(buffer, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		log.Error("jwt.validate: ", err.Error())
		return false
	}
	if !res.Valid {
		log.Error("jwt.validate: token invalid")
		return false
	}
	return true
}

func GetJwtToken(ctx echo.Context) []string {
	token := ctx.Request().Header["Authorization"][0]
	token = strings.Split(token, " ")[1]
	return strings.Split(token, ".")
}
