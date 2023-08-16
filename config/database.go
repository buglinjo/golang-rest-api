package config

import (
	"fmt"
	"github.com/buglinjo/golang-rest-api/migrations"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var DB *gorm.DB

type connection struct {
	Driver   string
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func getConnection() connection {
	return connection{
		Driver:   os.Getenv("DB_CONNECTION"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}

func InitDB() error {
	con := getConnection()
	dsn := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		con.Username,
		con.Password,
		con.Host,
		con.Port,
		con.Database,
	)

	db, err := gorm.Open(con.Driver, dsn)
	if err != nil {
		return err
	}

	DB = db

	migrations.AutoMigrate(DB)

	return nil
}

func CloseDBConnection() {
	err := DB.Close()
	if err != nil {
		log.Fatal("Error closing DB connection.")
	}
}
