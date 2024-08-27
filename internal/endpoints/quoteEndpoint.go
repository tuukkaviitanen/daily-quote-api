package endpoints

import (
	"daily_quote_api/internal/models"

	"github.com/gin-gonic/gin"
)

func GetQuoteEndpoint(context *gin.Context) {
	quote := models.Quote{Title: "Quote of the day", Quote: "This is a quote", Author: "me"}

	context.JSON(200, quote)
}
