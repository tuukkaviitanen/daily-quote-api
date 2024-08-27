package router

import (
	"daily_quote_api/internal/endpoints"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	quotes := router.Group("/quote")
	{
		quotes.GET("/", endpoints.GetQuoteEndpoint)
	}

	return router
}
