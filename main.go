package main

import (
	"log"
	"os"

	"github.com/buglinjo/golang-rest-api/app/routes"
	"github.com/buglinjo/golang-rest-api/config"
	"github.com/buglinjo/golang-rest-api/migrations"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	setupRouter(setupDB())
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func setupDB() *gorm.DB {
	db, err := config.Gorm()
	if err != nil {
		log.Fatal(err)
	}

	migrations.AutoMigrate(db)
	// defer db.Close()

	return db
}

func setupRouter(db *gorm.DB) {
	r := routes.Setup(db)
	r.Run(":" + os.Getenv("PORT"))
}
