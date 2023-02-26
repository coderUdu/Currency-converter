package main

import (
	"os"
	

	"github.com/gin-gonic/gin"
	"github.com/currencyconverter/routers"
)

func main() {
	port := os.Getenv("PORT")

	if port == " "{
		port ="8080"
	}

	router :=gin.New()
	router.Use(gin.Logger())

	routers.Authroutes(router)
}