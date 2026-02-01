package main

import (
	"net/http"
	"weather-service/internal/handlers"
)

func main() {
	http.HandleFunc("/weather", handlers.GetWeather)
	http.ListenAndServe(":8080", nil)
}
