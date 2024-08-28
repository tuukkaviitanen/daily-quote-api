package services

import (
	"daily-quote-api/internal/database"
	"daily-quote-api/internal/entities"
	"fmt"
)

func FetchQuote(index int) (entities.Quote, error) {
	var quote entities.Quote
	if err := database.Database.Find(&quote, index).Error; err != nil {
		fmt.Println("Database error occurred,", err)
		return quote, err
	}
	return quote, nil
}

func FetchQuoteCount() (int64, error) {
	var count int64
	if err := database.Database.Model(&entities.Quote{}).Count(&count).Error; err != nil {
		fmt.Println("Database error occurred,", err)
		return count, err
	}
	fmt.Println("Count:", count)
	return count, nil
}
