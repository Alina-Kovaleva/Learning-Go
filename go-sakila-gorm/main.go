package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Film struct {
	Title       string `gorm:"type:varchar(128)"`
	Description string `gorm:"type:text"`
	ReleaseYear int    `gorm:"type:year"`
	LanguageID  int    `gorm:"type:tinyint"`
	Length      int    `gorm:"type:smallint"`
}

func (Film) TableName() string {
	return "film"
}
func Connect() (*gorm.DB, error) {
	password := os.Getenv("DB_PASSWORD")
	dsn := "root:" + password + "@tcp(localhost:3306)/sakila"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetFilms(db *gorm.DB) ([]Film, error) {
	var films []Film
	err := db.Where("title LIKE ?", "%FAMILY%").Find(&films).Error
	// err := db.Where("title LIKE ?", "%Sakila%").Find(&films).Error
	if err != nil {
		return nil, err
	}
	return films, nil
}

func main() {
	// password := os.Getenv("DB_PASSWORD")
	// db, err := gorm.Open(mysql.Open("root:"+password+"@tcp(localhost:3306)/sakila"), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }

	// db.AutoMigrate(&Film{})
	// var film Film
	// // db.Find(&film, "title LIKE ?", "%FAMILY%")
	// db.Find(&film, "title LIKE ?", "%Sakila%")

	// log.Println(film)

	// db.Create(&Film{
	// 	Title:       "Sakila DB",
	// 	Description: "Test sakila db create film",
	// 	ReleaseYear: 2023,
	// 	LanguageID:  1,
	// 	Length:      120,
	// })

	db, err := Connect()
	if err != nil {
		panic(err)
	}

	films, err := GetFilms(db)
	if err != nil {
		panic(err)
	}

	for _, film := range films {
		fmt.Println(film.Title, film.Description)
	}

	// db.Create(&Film{
	// 	Title:       "Sakila FAMILY",
	// 	Description: "Test sakila db create new film with FAMILY in title",
	// 	ReleaseYear: 2023,
	// 	LanguageID:  1,
	// 	Length:      120,
	// })
}
