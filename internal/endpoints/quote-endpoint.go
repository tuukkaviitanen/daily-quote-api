package endpoints

import (
	"daily-quote-api/internal/enums"
	"daily-quote-api/internal/models"
	"daily-quote-api/internal/services"
	"daily-quote-api/internal/utils"
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
)

func GetQuoteEndpoint(context *gin.Context) {
	timeIntervalString := context.DefaultQuery("time_interval", string(enums.DAILY))

	timeInterval := enums.TimeInterval(timeIntervalString)

	if !timeInterval.IsValidTimeInterval() {
		context.JSON(400, gin.H{"error": "Invalid time_interval"})
		return
	}

	quotesCount, err := services.FetchQuoteCount()
	if err != nil {
		context.JSON(500, gin.H{"error": "Unexpected error occurred"})
		return
	}

	todayEpoch, unit := utils.IntervalToEpoch(timeInterval)

	seed := int64(todayEpoch)
	randomGenerator := rand.New(rand.NewSource(seed))
	randomNumber := randomGenerator.Intn(int(quotesCount))

	quote, err := services.FetchQuote(randomNumber)
	if err != nil {
		context.JSON(500, gin.H{"error": "Unexpected error occurred"})
		return
	}

	title := fmt.Sprintf("Quote of the %v", unit)

	returnBody := models.Quote{Title: title, Quote: quote.Quote, Author: quote.Author}

	context.JSON(200, returnBody)
}
