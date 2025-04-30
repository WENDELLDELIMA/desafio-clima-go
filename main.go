package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/weather/", WeatherHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
