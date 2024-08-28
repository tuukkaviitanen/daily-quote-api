package endpoints

import (
	"daily-quote-api/internal/enums"
	"daily-quote-api/internal/models"
	"daily-quote-api/internal/services"
	"daily-quote-api/internal/utils"
	"fmt"
	"math/rand"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetQuoteEndpoint(context *gin.Context) {
	unitOfTimeString := strings.ToLower(context.DefaultQuery("of-the", string(enums.DAY)))

	unitOfTime := enums.UnitOfTime(unitOfTimeString)

	if !unitOfTime.IsValidUnitOfTime() {
		context.JSON(400, gin.H{"error": fmt.Sprintf("Invalid unit of time: '%v'", unitOfTimeString)})
		return
	}

	quotesCount, err := services.FetchQuoteCount()
	if err != nil {
		context.JSON(500, gin.H{"error": "Unexpected error occurred"})
		return
	}

	todayEpoch := utils.UnitOfTimeToEpoch(unitOfTime)

	seed := int64(todayEpoch)
	randomGenerator := rand.New(rand.NewSource(seed))
	randomNumber := randomGenerator.Intn(int(quotesCount))

	quote, err := services.FetchQuote(randomNumber)
	if err != nil {
		context.JSON(500, gin.H{"error": "Unexpected error occurred"})
		return
	}

	title := fmt.Sprintf("Quote of the %v", unitOfTimeString)

	returnBody := models.Quote{Title: title, Quote: quote.Quote, Author: quote.Author}

	context.JSON(200, returnBody)
}
