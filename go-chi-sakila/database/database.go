package database

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	password := os.Getenv("DB_PASSWORD")
	conn := ("root:" + password + "@tcp(localhost:3306)/sakila")

	db, err := gorm.Open(
		mysql.Open(conn+"?parseTime=true"),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}

	DB = db
}
