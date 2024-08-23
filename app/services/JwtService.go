package services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type ClaimsArgs struct {
	LastIpAddress string
}

type CustomClaims struct {
	LastIpAddress string `json:"last_ip_address"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(claims_args ClaimsArgs) (string, string) {
	accessTokenId := uuid.NewString()

	claims := CustomClaims{
		LastIpAddress: claims_args.LastIpAddress,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Minute * 30)},
			Issuer:    "golang-auth",
			ID:        accessTokenId,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	jwtAccessToken, err := token.SignedString([]byte(os.Getenv("mySigningKey")))
	refreshToken := GenerateRefreshToken(accessTokenId)

	fmt.Println(claims)
	fmt.Println(jwtAccessToken, err, token)

	return jwtAccessToken, refreshToken
}

func GenerateRefreshToken(accessTokenId string) string {
	refreshToken, err := bcrypt.GenerateFromPassword([]byte(accessTokenId), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(refreshToken)
}
