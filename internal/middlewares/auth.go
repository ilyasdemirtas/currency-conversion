package middlewares

import (
	token "arf/currency-conversion/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
