package routes

import (
	"authentication-modules/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	v1 := r.Group("/api/v1")
	// v1Auth := r.Group("/api/v1/auth", UserAuthMiddleware())

	v1.GET("/", defaultHandler)
	return r
}

func defaultHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Default Page"})
}
