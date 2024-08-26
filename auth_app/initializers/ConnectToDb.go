package initializers

import (
	"auth_app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	migrateErr := DB.AutoMigrate(&models.User{})

	if err != nil {
		panic("database connection failed")
	}

	if migrateErr != nil {
		panic(migrateErr)
	}

}
