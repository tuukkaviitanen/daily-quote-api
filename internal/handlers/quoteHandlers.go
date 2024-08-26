package handlers

import (
	"fmt"
	"net/http"
)

func GetQuoteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
