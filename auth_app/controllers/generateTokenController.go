package controllers

import (
	"auth_app/initializers"
	"auth_app/models"
	"auth_app/services"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var UserFields struct {
	ID           uuid.UUID `json:"user_id" gorm:"id"`
	RefreshToken string    `json:"refresh_token"`
	AccessToken  string    `json:"access_token"`
}

func GenerateTokens(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := c.Bind(&UserFields); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	var user models.User
	if response := initializers.DB.Take(&user, "ID = ?", UserFields.ID); response.Error != nil {
		c.JSON(400, gin.H{
			"error": "User not found",
		})

		return
	}

	accessToken, refreshToken := services.GenerateAccessToken(UserFields.ID, getClientIpAddress(c))

	user.LastIpAddress = getClientIpAddress(c)
	user.RefreshToken = refreshToken
	initializers.DB.Save(&user)

	c.JSON(200, gin.H{
		"access_token":  accessToken,
		"refresh_token": services.EncodeRefreshToken(refreshToken),
	})
}

func RefreshTokens(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := c.Bind(&UserFields); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	UserFields.ID, _, _ = services.DecodeRefreshToken(UserFields.RefreshToken)

	var user models.User
	if response := initializers.DB.Take(&user, "ID = ?", UserFields.ID); response.Error != nil {
		c.JSON(400, gin.H{
			"error": "Invalid refresh token",
		})

		return
	}

	refreshTokenData, _ := base64.StdEncoding.DecodeString(UserFields.RefreshToken)
	err := bcrypt.CompareHashAndPassword([]byte(user.RefreshToken), refreshTokenData)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid refresh token",
		})

		return
	}

	accessToken, refreshToken, err := services.ValidateRefreshToken(UserFields.AccessToken, UserFields.RefreshToken, getClientIpAddress(c))

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	user.RefreshToken = refreshToken
	initializers.DB.Save(&user)

	c.JSON(200, gin.H{
		"access_token":  accessToken,
		"refresh_token": services.EncodeRefreshToken(refreshToken),
	})
}

func getClientIpAddress(c *gin.Context) string {
	ipAddress := c.Request.Header.Get("X-Real-Ip")
	if ipAddress == "" {
		ipAddress = c.Request.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress = c.Request.RemoteAddr
	}

	return ipAddress
}
