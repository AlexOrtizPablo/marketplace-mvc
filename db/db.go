package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"marketplace-mvc/model"
)

var Instance = connectDB()

func connectDB() *gorm.DB {
	dbUser := "root"
	dbPass := "secret"
	dbHost := "localhost"
	dbName := "marketplace_mvc"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Aruba", dbHost, dbUser, dbPass, dbName)
	db, errorDB := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errorDB != nil {
		panic("Failed to connect postgres database")
	}

	return db
}

func ExecuteMigrations() error {
	DB := Instance
	err := DB.AutoMigrate(&model.Address{}, &model.Payment{}, &model.Product{}, &model.Purchase{},
		&model.Question{}, &model.Rating{}, &model.Shipping{}, &model.ShoppingCart{}, &model.User{})
	if err != nil {
		return err
	}

	return nil
}
