package main

import (
	"daily_quote_api/internal/routers"
	"fmt"
	"net/http"
)

func main() {
	server := http.NewServeMux()

	quoteRouter := routers.GetQuoteRouter()

	server.Handle("/quote", quoteRouter)

	if err := http.ListenAndServe("localhost:8080", server); err != nil {
		fmt.Println(err.Error())
	}
}
