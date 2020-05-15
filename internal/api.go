package api

import (
	"github.com/MiXALK/go-API/internal/controllers"
	"github.com/MiXALK/go-API/internal/utils"
	"github.com/joho/godotenv"
	"os"
)

var server = controllers.APIServer{}

func init() {
	errLoad := godotenv.Load()
	utils.ErrorCheck(errLoad)
}

func Run() {
	server.Initialize(
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))
	server.Run(os.Getenv("API_PORT"))
}
