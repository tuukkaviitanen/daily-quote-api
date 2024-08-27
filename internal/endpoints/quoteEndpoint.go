package endpoints

import (
	"daily_quote_api/internal/services"

	"github.com/gin-gonic/gin"
)

func GetQuoteEndpoint(context *gin.Context) {
	quote := services.FetchQuote(1)
	//returnBody := models.Quote{Title: "Quote of the day", Quote: "This is a quote", Author: "me"}

	context.JSON(200, quote)
}
