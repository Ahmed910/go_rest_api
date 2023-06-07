package route

import "github.com/gin-gonic/gin"

func HandleRoutes(router *gin.Engine) *gin.Engine {
	apiRouter := router.Group("/api")
	videoRouter := apiRouter.Group("/videos")
	VideoRoutes(videoRouter)
	return router
}
