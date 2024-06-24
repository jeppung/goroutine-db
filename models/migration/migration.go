package main

import (
	"fmt"

	"github.com/jeppung/goroutine-db/database"
	"github.com/jeppung/goroutine-db/models"
)

func init() {
	database.ConnectToDB()
}

func main() {

	if err := database.Client.AutoMigrate(&models.User{}); err != nil {
		fmt.Println("Error migrating")
	}else{
		fmt.Println("Success migrating")
	}
	
}