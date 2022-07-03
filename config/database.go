package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"5432",
		"rakaadli",
		"qawsedrf",
		"postgres",
	)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}

	log.Default().Println("Connection to Database is Successfull")

	// db.AutoMigrate(models.Item{}, models.Order{})
	return db
}
