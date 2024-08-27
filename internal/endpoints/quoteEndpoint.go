package endpoints

import (
	"daily_quote_api/internal/models"
	"daily_quote_api/internal/services"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func GetQuoteEndpoint(context *gin.Context) {
	quotesCount := services.FetchQuoteCount()
	now := time.Now()
	todayEpoch := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()

	seed := int64(todayEpoch)
	randomGenerator := rand.New(rand.NewSource(seed))
	randomNumber := randomGenerator.Intn(int(quotesCount))

	quote := services.FetchQuote(randomNumber)

	returnBody := models.Quote{Title: "Quote of the day", Quote: quote.Quote, Author: quote.Author}

	context.JSON(200, returnBody)
}
