package main

import (
	"fmt"
	"github.com/buglinjo/golang-rest-api/app/routes"
	"github.com/buglinjo/golang-rest-api/config"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	envFile = ".env"
)

func main() {
	err := Run()
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}

func Run() error {
	err := loadEnv(envFile)
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	err = setupDB()
	if err != nil {
		return fmt.Errorf("error setting up database: %v", err)
	}

	defer closeDB()

	err = setupServer()
	if err != nil {
		return fmt.Errorf("error setting up server: %v", err)
	}

	return nil
}

func loadEnv(filename string) error {
	return godotenv.Load(filename)
}

func setupDB() error {
	return config.InitDB()
}

func closeDB() {
	config.CloseDBConnection()
}

func setupServer() error {
	r := routes.Setup()

	return r.Run(":" + os.Getenv("PORT"))
}
