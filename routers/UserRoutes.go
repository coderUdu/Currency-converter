package routers

import (
	"github.com/currencyconverter/controllers"
	"github.com/currencyconverter/middleware"
	"github.com/gin-gonic/gin"
)



func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())

	incomingRoutes.GET("/users/GetBalance", controllers.GetBalanceHistory())
	incomingRoutes.GET("/users", controllers.ConvertCurrency())
	
}