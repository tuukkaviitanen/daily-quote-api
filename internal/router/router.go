package router

import (
	"daily-quote-api/internal/endpoints"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	quotes := router.Group("/quote")
	{
		quotes.GET("/", endpoints.GetQuoteEndpoint)
	}

	// Redirect / and /swagger to /swagger/
	router.GET("/swagger", func(c *gin.Context) {
		c.Redirect(301, "/swagger/")
	})
	router.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/swagger/")
	})

	router.StaticFile("/swagger.yaml", "./api/openapi.yaml")
	router.Static("/swagger", "./api/swagger-ui")

	return router
}
