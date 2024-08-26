package handlers

import (
	"daily_quote_api/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetQuoteHandler(w http.ResponseWriter, r *http.Request) {
	quote := models.Quote{Title: "Quote of the day", Quote: "This is a quote", Author: "me"}

	jsonBytes, err := json.Marshal(quote)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "Unexpected error occurred")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonBytes)
}
