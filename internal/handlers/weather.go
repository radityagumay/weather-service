package handlers

import (
	"encoding/json"
	"net/http"
	"weather-service/internal/data"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	weatherData, err := data.LoadWeatherData()
	if err != nil {
		http.Error(w, "Unable to load weather data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherData)
}
