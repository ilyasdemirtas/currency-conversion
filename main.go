package main

import (
	"arf/currency-conversion/app"
	"arf/currency-conversion/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	app.LoadEnv()
	app.DbConn()

	routes.GetAPIRoutes(r)
	routes.GetSwaggerRoute(r)

	r.Run(":8080")
}
