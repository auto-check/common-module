package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type Claims struct {
	StudentID int64
	jwt.StandardClaims
}

var (
	atExpirationTime = time.Now().Add(time.Hour * 1).Unix()
	rtExpirationTime = time.Now().Add(time.Hour * 24 * 14).Unix()
	accessSecret = os.Getenv("ACCESS_SECRET")
	refreshSecret = os.Getenv("REFRESH_SECRET")
)