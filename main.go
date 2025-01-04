package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pchihieuu/url-shortner.git/api/database"
	"github.com/pchihieuu/url-shortner.git/api/routes"
)

func main() {
	err := godotenv.Load()
	if err!= nil {
        fmt.Println(err)
    }

	database.InitializeClient()
	
	router := gin.Default()
	setupRouters(router) 
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(router.Run(":" + port))
}

func setupRouters(router *gin.Engine) {
	router.POST("/api/v1", routes.ShortenURL)
	router.GET("/api/v1/:shortID", routes.GetByShortID)
	router.DELETE("/api/v1/:shortID", routes.DeleteURL)
	router.PUT("/api/v1/:shortID", routes.EditURL)
	router.POST("/api/v1/addTag", routes.AddTag)
}