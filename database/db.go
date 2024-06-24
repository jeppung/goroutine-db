package database

import (
	"fmt"
	"os"

	"github.com/jeppung/goroutine-db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Client *gorm.DB

func ConnectToDB() {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", 
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	if db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		panic("Error connect to db")
	}else{
		Client = db
	}
}

func Migration() {
	err := Client.AutoMigrate(&models.User{})
	if err != nil {
		panic("Error migration")
	}
}