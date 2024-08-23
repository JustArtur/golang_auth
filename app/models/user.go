package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`

	LastIpAddress string
	RefreshToken  string
	CreatedAt     time.Time // Automatically managed by GORM for creation time
	UpdatedAt     time.Time // Automatically managed by GORM for update time
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	hashedRefreshToken, err := generateHash(u.RefreshToken)

	u.RefreshToken = hashedRefreshToken
	return
}

func generateHash(object string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(object), bcrypt.DefaultCost)

	return string(hashed), err
}
