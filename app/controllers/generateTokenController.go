package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_jwt_auth/app/initializers"
	"golang_jwt_auth/app/models"
	"golang_jwt_auth/app/services"
)

func GenerateTokens(c *gin.Context) {
	var User struct {
		ID string `json:"guid" gorm:"guid"`
	}

	if err := c.Request.ParseForm(); err != nil {
		panic(err)
	}

	if err := c.Bind(&User); err != nil {
		c.JSON(200, gin.H{
			"error": err,
		})

		return
	}

	var user models.User

	fmt.Println(initializers.DB.First(&user, "ID = ?", userID))

	fmt.Println(user.CreatedAt)

	accessToken, refreshToken := services.GenerateAccessToken(
		struct {
			UserID        string
			LastIpAddress string
		}{
		UserID:        User.ID,
		LastIpAddress: c.Request.RemoteAddrg
		}
	)

	//User.RefreshToken = refreshToken

	c.JSON(200, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
