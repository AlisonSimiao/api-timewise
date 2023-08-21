package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB = nil

// get connection to database with env
func Connect() error {
	DB_PASS := os.Getenv("DB_PASS")
	DB_USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")

	db, err := gorm.Open(postgres.Open("user="+DB_USER+" password="+DB_PASS+" host="+DB_HOST+" port="+DB_PORT+" dbname="+DB_NAME), &gorm.Config{
		PrepareStmt: false,
	})

	if err != nil {
		fmt.Printf("ERR> database connection failed %s", err)
		return err
	}

	database = db

	return nil
}

// get connection to database with string
func ConnectWithString(conn string) error {
	DB_CONNECTION := conn
	db, err := gorm.Open(postgres.Open("user=postgres password="+DB_CONNECTION+" host=db.yoegocppoqdjwvluebjn.supabase.co port=5432 dbname=postgres"), &gorm.Config{
		PrepareStmt:            false,
		SkipDefaultTransaction: true,
	})

	if err != nil {
		fmt.Printf("ERR> database connection failed: %s", DB_CONNECTION)
		return err
	}

	database = db

	return nil
}

func GetDatabase() *gorm.DB {
	return database
}
