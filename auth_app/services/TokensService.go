package services

import (
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	"strconv"
	"strings"
	"time"
)

type CustomClaims struct {
	LastIpAddress string `json:"last_ip_address"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(UserID uuid.UUID, LastIpAddress string) (string, string) {

	claims := CustomClaims{
		LastIpAddress: LastIpAddress,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Minute * 30)},
			Issuer:    "golang-auth",
			Subject:   UserID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	accessToken, _ := token.SignedString([]byte(os.Getenv("mySigningKey")))
	refreshToken := GenerateRefreshToken(UserID.String(), LastIpAddress)

	return accessToken, refreshToken
}

func GenerateRefreshToken(UserID string, LastIpAddress string) string {
	refreshToken := fmt.Sprintf("ID:%s;IP:%s;exp:%d", UserID, LastIpAddress, time.Now().Add(24*time.Hour).Unix())

	return refreshToken
}

func EncodeRefreshToken(refreshToken string) string {
	return base64.StdEncoding.EncodeToString([]byte(refreshToken))
}

func ValidateRefreshToken(accessToken string, refreshToken string, currentIpAddress string) (string, string, error) {
	userID, userIP, exp := DecodeRefreshToken(refreshToken)
	accesыTokenClaims, err := DecodeAccessToken(accessToken)

	if err != nil {
		return "", "", err
	}

	accesTokenSubject, err := accesыTokenClaims.GetSubject()

	if err != nil {
		return "", "", err
	}

	accessTokenUserID, err := uuid.Parse(accesTokenSubject)
	if err != nil {
		return "", "", err
	}

	if accessTokenUserID != userID {
		return "", "", fmt.Errorf("invalid refresh token")
	}

	if userIP != currentIpAddress {
		SendWarningEmail(userID)
	}

	expirationTime, err := strconv.ParseInt(exp, 10, 64)
	expTime := time.Unix(expirationTime, 0)

	if err != nil {
		return "", "", err
	} else if expTime.Before(time.Now()) {
		return "", "", fmt.Errorf("token expired")
	}

	newAccessToken, newRefreshToken := GenerateAccessToken(accessTokenUserID, userIP)

	return newAccessToken, newRefreshToken, nil
}

func DecodeAccessToken(accessToken string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("mySigningKey")), nil
	})

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*CustomClaims)

	return claims, nil
}

func DecodeRefreshToken(refreshToken string) (uuid.UUID, string, string) {
	refreshTokenData, _ := base64.StdEncoding.DecodeString(refreshToken)

	tokenData := strings.Split(string(refreshTokenData), ";")

	userID, _ := uuid.Parse(strings.Split(tokenData[0], ":")[1])

	return userID, strings.Split(tokenData[1], ":")[1], strings.Split(tokenData[2], ":")[1]
}
