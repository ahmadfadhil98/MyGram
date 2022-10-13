package infrastructure

import (
	"MyGram/domain"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"
)

var Database SqlHandler

type SqlHandler struct {
	DB *gorm.DB
}

func (dbs *SqlHandler) DBInit() error {
	err := godotenv.Load("../MyGram/app/.env")
	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"), os.Getenv("DB_PORT"))
	fmt.Println(dataSourceName)
	db, err := gorm.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(domain.User{}, domain.Photo{}, domain.SocialMedia{}, domain.Comment{})
	dbs.DB = db
	return err
}
