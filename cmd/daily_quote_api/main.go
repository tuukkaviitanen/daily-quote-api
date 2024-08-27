package main

import (
	"daily_quote_api/internal/database"
	"daily_quote_api/internal/router"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8080"
	}

	port, err := strconv.ParseInt(portString, 0, 64)
	if err != nil {
		log.Fatalln("PORT is invalid, Error:", err)
	}

	database.ConnectToDatabase()

	router := router.GetRouter()

	address := fmt.Sprintf(":%v", port)

	if err := router.Run(address); err != nil {
		log.Fatalln("Error occurred while running the server, Error", err)
	}
}
