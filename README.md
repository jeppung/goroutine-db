# Goroutine DB

Goroutine concurrency demo inserting **1.000.000 data**  with 
- **Excelize** (Reading excel data)
- **PostgreSQL** (Database)
- **Gorm** (ORM)

You are going to see how fast it takes to insert 1.000.000 data to database

## !! Important !!
- timer start after reading excel data
- result may vary with different hardware specification

## How to run
- clone this repo
- create postgres db named -> goroutine-db
- rename .env.example -> .env
- input your postgres db config in .env
    - example:
        - DB_HOST=localhost
        - DB_USER=postgres
        - DB_PASSWORD=postgres
        - DB_NAME=goroutine-db
        - DB_PORT=5432
- go mod tidy
- run code
    - Code with goroutine -> go run main.go
    - Code without goroutine -> go run single-worker/main.go
