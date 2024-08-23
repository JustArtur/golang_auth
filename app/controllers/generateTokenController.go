package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang_jwt_auth/app/initializers"
	"golang_jwt_auth/app/models"
	"golang_jwt_auth/app/services"
)

func GenerateTokens(c *gin.Context) {
	var UserFields struct {
		ID uuid.UUID `json:"user_id" gorm:"id"`
	}

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
	if response := initializers.DB.First(&user, "ID = ?", UserFields.ID); response.Error != nil {
		c.JSON(400, gin.H{
			"error": "User not found",
		})

		return
	}

	accessToken, refreshToken := services.GenerateAccessToken(UserFields.ID.String(), getClientIpAddress(c))

	user.LastIpAddress = getClientIpAddress(c)
	user.RefreshToken = refreshToken
	fmt.Println(initializers.DB.Save(&user))

	c.JSON(200, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
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
