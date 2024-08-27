package services

import (
	"daily_quote_api/internal/database"
	"daily_quote_api/internal/entities"
	"fmt"
)

func FetchQuote(index int) entities.Quote {
	var quote entities.Quote
	database.Database.Find(&quote, index)
	return quote
}

func FetchQuoteCount() int64 {
	var count int64
	database.Database.Model(&entities.Quote{}).Count(&count)
	fmt.Println("Count:", count)
	return count
}
