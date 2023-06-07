package main

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/route"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()
	route.HandleRoutes(server)
	server.Run(":3000")
}
