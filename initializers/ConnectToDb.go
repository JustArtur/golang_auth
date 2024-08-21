package initializers

import (
	"golang_jwt_auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := "host=localhost dbname=golang_auth_delopment port=5432"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	migrate_err := DB.AutoMigrate(&models.User{})

	if err != nil {
		panic("database connection failed")
	}

	if migrate_err != nil {
		panic(migrate_err)
	}
}
