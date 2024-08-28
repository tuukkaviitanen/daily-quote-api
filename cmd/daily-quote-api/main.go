package main

import (
	"daily-quote-api/internal/database"
	"daily-quote-api/internal/router"
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
