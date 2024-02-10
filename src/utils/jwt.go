package utils

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mehrdadjalili/facegram_common/pkg/derrors"
	"github.com/mehrdadjalili/facegram_common/pkg/translator/messages"
)

func GenerateJwtToken(key, data string, exp int64) (string, error) {
	password := []byte(key)
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["data"] = data
	claims["exp"] = exp

	tokenString, err := token.SignedString(password)

	if err != nil {
		return "", derrors.InternalError()
	}

	return tokenString, nil
}

func ParseJwtToken(signingKey, userToken string) (string, error) {
	password := []byte(signingKey)
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(userToken, claims, func(token *jwt.Token) (interface{}, error) {
		return password, nil
	})
	if err != nil || !token.Valid {
		return "", derrors.New(derrors.StatusUnauthorized, messages.Unauthorized)
	}

	data := fmt.Sprintf("%v", claims["data"])

	//if int64(exp) < time.Now().Unix() {
	//	return "", derrors.New(derrors.StatusUnauthorized, messages.Unauthorized)
	//}
	return data, nil
}
