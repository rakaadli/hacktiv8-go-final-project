package config

import (
	"fmt"
	"hacktiv8-final-project/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname =%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	log.Default().Println("Connection to Database is Successfull")

	db.AutoMigrate(models.User{})
	return db
}
