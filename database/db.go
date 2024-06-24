package database

import (
	"github.com/jeppung/goroutine-db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Client *gorm.DB

func ConnectToDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=goroutine-db port=5432 sslmode=disable TimeZone=Asia/Jakarta"
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