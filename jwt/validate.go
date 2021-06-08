package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

func Validate(token string) (int64, error){
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token)(interface{}, error){
		return []byte(accessSecret), nil
	})

	if err != nil {
		return -1, err
	}

	return claims.StudentID, nil
}
