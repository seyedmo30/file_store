package repository

import (
	"fmt"
	"os"
	"retrieval/entity"
	"retrieval/pkg/encryption"
	"retrieval/pkg/logs"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

func GetDBInstance() (*gorm.DB, error) {
	var err error

	dbOnce.Do(func() {
		POSTGRES_GATE_WAY_HOST := os.Getenv("POSTGRES_GATE_WAY_HOST")
		POSTGRES_GATE_WAY_PORT := os.Getenv("POSTGRES_GATE_WAY_PORT")
		POSTGRES_GATE_WAY_USER := os.Getenv("POSTGRES_GATE_WAY_USER")
		POSTGRES_GATE_WAY_PASSWORD := os.Getenv("POSTGRES_GATE_WAY_PASSWORD")
		POSTGRES_GATE_WAY_DB := os.Getenv("POSTGRES_GATE_WAY_DB")
		dsn := "host=" + POSTGRES_GATE_WAY_HOST + " user=" + POSTGRES_GATE_WAY_USER + " password=" + POSTGRES_GATE_WAY_PASSWORD + " dbname=" + POSTGRES_GATE_WAY_DB + " port=" + POSTGRES_GATE_WAY_PORT + " sslmode=disable"
		dbInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("Error connecting to database:", err)
			return
		}
		fmt.Println("Database connection established.")

		dbInstance.AutoMigrate(&entity.User{})
	})

	return dbInstance, err
}

func CreateUser(username string, password string) error {

	db, err := GetDBInstance()
	if err != nil {
		logs.Connect().Error("connection database refuse" + err.Error())

		return err
	}
	hashPassword, err := encryption.GetHash(password)

	if err != nil {

		logs.Connect().Error("hash Password error :" + err.Error())
		return err

	}
	user := entity.User{
		Username:     username,
		PasswordHash: hashPassword,
	}
	result := db.Create(&user)

	if result.Error != nil {

		logs.Connect().Error("Create user error :" + err.Error())
		return result.Error

	}
	logs.Connect().Info("User created:" + user.Username)

	return nil
}

func CheckPassword(username string, password string) (bool, error) {
	db, err := GetDBInstance()
	if err != nil {
		logs.Connect().Error("connection database refuse" + err.Error())

		return false, err
	}
	var fetchedUser entity.User
	db.First(&fetchedUser, "username = ?", username)

	if fetchedUser.Username == "" {
		logs.Connect().Info("user not found")

	}
	if isCorrect := encryption.CheckPassword(password, fetchedUser.PasswordHash); isCorrect {

		logs.Connect().Info("password is correct")
		return true, nil
	} else {

		logs.Connect().Info("password is incorrect")
		return false, nil
	}
}
