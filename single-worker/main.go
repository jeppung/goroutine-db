package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/jeppung/goroutine-db/database"
	"github.com/jeppung/goroutine-db/models"
	"github.com/joho/godotenv"
	"github.com/xuri/excelize/v2"
)

func init() {
	godotenv.Load()
	database.ConnectToDB()
}

func main() {
	fmt.Println("Reading excel data...")
	excelData, err := readingExcel()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
	start := time.Now()
	for _, data := range excelData {
		user := &models.User{
			Username: data[0],
			Firstname: data[1],
			Lastname: data[2],
			Email: data[3],
		}
		
		AddDataToDb(user)
	}

	fmt.Println("\nData insertion complete...")
	fmt.Println("Took", time.Since(start))
}

func readingExcel() ([][]string, error) {
	f, err := excelize.OpenFile("./user_data.xlsx")
	defer f.Close()

	if err != nil {
		return nil, errors.New("Error reading excel")
	}

	datas, err := f.GetRows("Sheet1")
	if err != nil {
		return nil, errors.New("Error reading rows in sheetbook")
	}

	return datas[1:], nil
}

func AddDataToDb(user *models.User) {
	count := 0

	for {
		res := database.Client.Create(user)

		if res.Error == nil {
			fmt.Printf("Worker(1) done working on inserting %v\n", user.Username)
			break
		}

		if count == 3 {
			fmt.Println("Worker(1) fail on inserting %v", user.Username)
			break
		}

		fmt.Printf("Worker(1) retrying on inserting %v\n", user.Username)
		count++
	}
}