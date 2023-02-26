package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/currencyconverter/controllers"
) 

func Authroutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("users/signup", controllers.SignUp())

	incomingRoutes.POST("users/login", controllers.Login())

	
}