package main

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/jeppung/goroutine-db/database"
	"github.com/jeppung/goroutine-db/models"
	"github.com/joho/godotenv"
	"github.com/xuri/excelize/v2"
)

func init() {
	godotenv.Load() // Loading env variable
	database.ConnectToDB()
	database.Migration()
}

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	jobChannel := make(chan *models.User, 100)

	// Reading excel
	fmt.Println("Reading excel data...")
	excelData, err := readingExcel()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
	// Start working with data from excel
	go addWorker(50, jobChannel, &wg)
	addJob(excelData, jobChannel, &wg)
	
	// Wait for all goroutine to complete job task
	wg.Wait()

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

func addJob(datas [][]string, ch chan <- *models.User, wg *sync.WaitGroup) {
	for _, data := range datas {
		user := &models.User{
			Username: data[0],
			Firstname: data[1],
			Lastname: data[2],
			Email: data[3],
		}
		
		wg.Add(1)
		ch <- user
	}
	close(ch)
}

func addWorker(workers uint, ch <- chan  *models.User, wg *sync.WaitGroup) {
	for i:=0; i<int(workers); i++ {
		go func() {
			for data := range ch {
				fmt.Printf("Worker(%d) working on inserting %v\n", i, data.Username)
				AddDataToDb(data, i, wg)
			}
        }()
	}
}

func AddDataToDb(user *models.User, workerId int, wg *sync.WaitGroup) {
	for {
		res := database.Client.Create(user)

		if res.Error == nil {
			wg.Done()
			fmt.Printf("Worker(%d) done working on inserting %v\n", workerId, user.Username)
			break
		}else{
			fmt.Println("Error:", res.Error)
			os.Exit(1)
		}

		fmt.Printf("Worker(%d) retrying on inserting %v\n", workerId, user.Username)
	}
}