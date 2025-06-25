package main

import (
	"water/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Println("Error loading .env file")
	}
	ginMode := os.Getenv("GIN_MODE")
	gin.SetMode(ginMode)

	routes.Run()
}