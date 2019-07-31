package config

import (
	"os"
	"github.com/jinzhu/gorm"
)

type connection struct {
	Connection string
	Host       string
	Port       string
	Database   string
	Username   string
	Password   string
}

func mysql() connection {
	return connection{
		Connection: os.Getenv("DB_CONNECTION"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		Database:   os.Getenv("DB_DATABASE"),
		Username:   os.Getenv("DB_USERNAME"),
		Password:   os.Getenv("DB_PASSWORD"),
	}
}

func Gorm() (db *gorm.DB, err error) {
	return gorm.Open(
		mysql().Connection,
		mysql().Username+
			":"+
			mysql().Password+
			"@("+
			mysql().Host+
			":"+
			mysql().Port+
			")/"+
			mysql().Database+
			"?charset=utf8mb4&parseTime=True&loc=Local",
	)
}
