package middleware

import (
	"fmt"
	"net/http"

	helper "github.com/currencyconverter/Helper"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
	
	if clientToken == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("")})
		c.Abort()
		return
	}

	claims, err := helper.ValidateToken(clientToken)
	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()
		return
	}
	c.Set("email", claims.Email)
	c.Set("uid", claims.UID)
	c.Set("userType", claims.UserType)
	c.Next()
}
}
