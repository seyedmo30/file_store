package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db              *sql.DB
	once            sync.Once
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "test"
	dbPassword = "test"
	dbName     = "test"
)

func initializeDB() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() (*sql.DB) {
	once.Do(func() {
		initializeDB()
	})

	return db
}
