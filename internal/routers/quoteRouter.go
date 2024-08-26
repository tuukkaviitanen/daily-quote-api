package routers

import (
	"daily_quote_api/internal/handlers"
	"net/http"
)

func GetQuoteRouter() *http.ServeMux {
	quoteRouter := http.NewServeMux()

	quoteRouter.HandleFunc("GET /", handlers.GetQuoteHandler)

	return quoteRouter
}
