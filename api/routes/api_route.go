package routes

import (
	"arf/currency-conversion/api/internal/handlers"
	"arf/currency-conversion/api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func GetAPIRoutes(r *gin.Engine) {

	router := r.Group("/api")
	router.POST("/login", handlers.Login)

	router.Use(middlewares.VerifyToken())
	router.GET("/user-wallet-accounts", handlers.UserWalletAccounts)
	router.POST("/create-exchange-offer", handlers.CreateExchangeOffer)
	router.POST("/send-offer", handlers.CreateExchangeTransaction)

}
