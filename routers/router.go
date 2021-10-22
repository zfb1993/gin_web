package routers

import (
	"github.com/gin-gonic/gin"
	"gin_web/controllers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.GET("/register",controllers.RegisterGet)
	return router
}