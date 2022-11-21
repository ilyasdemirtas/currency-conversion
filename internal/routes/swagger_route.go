package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
)

func GetSwaggerRoute(r *gin.Engine) {

	r.StaticFileFS("/swagger.yml", "swagger.yml", http.Dir("swagger"))

	opts := middleware.RedocOpts{SpecURL: "/swagger.yml"}
	sh := middleware.Redoc(opts, nil)
	r.Handle("GET", "/docs", gin.WrapH(sh))
}
