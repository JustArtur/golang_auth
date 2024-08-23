package services

import (
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type CustomClaims struct {
	LastIpAddress string `json:"last_ip_address"`
	UserId        string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(UserID string, LastIpAddress string) (string, string) {

	claims := CustomClaims{
		LastIpAddress: LastIpAddress,
		UserId:        UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Minute * 30)},
			Issuer:    "golang-auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	accessToken, _ := token.SignedString([]byte(os.Getenv("mySigningKey")))
	refreshToken := GenerateRefreshToken(UserID, LastIpAddress)

	return accessToken, refreshToken
}

func GenerateRefreshToken(UserID string, LastIpAddress string) string {
	refreshTokenData := fmt.Sprintf("userID:%s;userIP:%s;exp:%d", UserID, LastIpAddress, time.Now().Add(24*time.Hour).Unix())

	refreshToken := base64.StdEncoding.EncodeToString([]byte(refreshTokenData))

	return refreshToken
}
