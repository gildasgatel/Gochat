package handler

import "github.com/gin-gonic/gin"

func Routes(h Handler) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", h.Home)

	return router
}
