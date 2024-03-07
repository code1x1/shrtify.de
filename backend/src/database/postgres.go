package database

import (
	"errors"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"shrtify.de/src/types"
)

func NewPostgres() *gorm.DB {
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")
	db, err := gorm.Open(postgres.Open("host="+host+" port="+port+"  user="+user+"  password="+password+"  dbname="+dbname+"  sslmode=disable TimeZone=Europe/Berlin"), &gorm.Config{})
	if err != nil {
		panic(errors.Join(errors.New("database connection failed"), err))
	}
	err = db.AutoMigrate(&types.ShortUrl{})
	if err != nil {
		panic(errors.Join(errors.New("migration failed"), err))
	}
	return db
}
