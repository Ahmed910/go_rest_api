package route

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/controller"
	"gilab.com/pragmaticreviews/golang-gin-poc/middleware"
	"github.com/gin-gonic/gin"
)

func VideoRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	//defer videoRepository.CloseDB()

	router.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth())

	router.GET("", controller.FindAll)

	router.POST("", controller.Save)

	router.PUT("/:id", controller.Update)

	router.DELETE("/:id", controller.Delete)
	return router
}
