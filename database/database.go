package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB               *sql.DB
	connectionString string = fmt.Sprintf("%v:%v@tcp(%v)/VETERINAYA", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_URL"))
)

func ConnectDB() error {
	db, err := sql.Open("mysql", connectionString)

	if err := db.Ping(); err != nil {
		return err
	}
	DB = db

	log.Println("Connected to the database")

	return err
}
