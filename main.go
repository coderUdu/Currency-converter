package main

import (
	"log"
	"os"

	"github.com/currencyconverter/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .enc file")
	}

	port := os.Getenv("PORT")

	if port == " "{
		port ="8080"
	}

	router :=gin.New()
	router.Use(gin.Logger())

	routers.Authroutes(router)
	routers.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context){
		c.JSON(200, gin.H{"success":"access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context){
		c.JSON(200, gin.H{"success":"access granted for api-2"})
	})
}