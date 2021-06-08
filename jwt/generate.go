package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"runtime/debug"
)

func GenerateToken(id int64) (string, string, error) {
	atClaims := &Claims{
		StudentID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: atExpirationTime,
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	signedAT, err := at.SignedString([]byte(accessSecret))
	if err != nil {
		log.Error(fmt.Sprintf("Error %s\n%s", err, debug.Stack()))
		return "", "", err
	}

	rtClaims := &Claims{
		StudentID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: rtExpirationTime,
		},
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	signedRT, err := rt.SignedString([]byte(refreshSecret))
	if err != nil {
		log.Error(fmt.Sprintf("Error %s\n%s", err, debug.Stack()))
		return "", "", err
	}

	return signedAT, signedRT, err
}