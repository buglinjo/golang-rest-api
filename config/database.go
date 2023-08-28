package config

import (
	"errors"
	"fmt"
	"github.com/buglinjo/golang-rest-api/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

type driver struct {
	Name string
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

func (c *connection) getDriver() (d driver, err error) {
	switch c.Driver {
	case "mysql":
		return driver{Name: "mysql"}, nil
	case "postgres":
		return driver{Name: "postgres"}, nil
	default:
		return driver{}, errors.New("unsupported DB driver")
	}
}

func (d *driver) open(con *connection) (db *gorm.DB, err error) {
	if d.Name == "mysql" {
		dns := fmt.Sprintf(
			"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			con.Username,
			con.Password,
			con.Host,
			con.Port,
			con.Database,
		)

		return gorm.Open(mysql.Open(dns), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else if d.Name == "postgres" {
		dns := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/New_York",
			con.Host,
			con.Username,
			con.Password,
			con.Database,
			con.Port,
		)

		return gorm.Open(postgres.Open(dns), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	}

	return nil, errors.New("could not open DB connection")
}

func InitDB() error {
	con := getConnection()
	driver, err := con.getDriver()
	if err != nil {
		log.Fatalf("unsupported database driver: %s", err.Error())
	}

	db, err := driver.open(&con)
	if err != nil {
		return err
	}

	DB = db

	migrations.AutoMigrate(db)

	return nil
}

func CloseDBConnection() {
	db, err := DB.DB()
	err = db.Close()
	if err != nil {
		log.Fatal("Error closing DB connection.")
	}
}
