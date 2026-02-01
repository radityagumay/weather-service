package data

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "weather-service/internal/models"
)

func LoadWeatherData() ([]models.Weather, error) {
    file, err := os.Open("static/weather.json")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    byteValue, err := ioutil.ReadAll(file)
    if err != nil {
        return nil, err
    }

    var result struct {
        Weather []models.Weather `json:"weather"`
    }
    err = json.Unmarshal(byteValue, &result)
    if err != nil {
        return nil, err
    }

    return result.Weather, nil
}