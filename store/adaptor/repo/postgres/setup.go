package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"store/pkg/logs"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

type Setup struct {
}

func NewPostgres() Setup {
	return Setup{}
}
func initializeDB() {
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		logs.Connect().Error(err.Error())
	}
	err = db.Ping()
	if err != nil {
		logs.Connect().Error(err.Error())
	}
}

func GetDB() *sql.DB {
	once.Do(func() {
		initializeDB()
	})

	return db
}
