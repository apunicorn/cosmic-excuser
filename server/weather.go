package main

import (
	"encoding/json"
	"net/http"
)

type WeatherResponse struct {
	Current struct {
		Temperature float64 `json:"temperature_2m"`
		Humidity    int     `json:"relative_humidity_2m"`
		WindSpeed   float64 `json:"wind_speed_10m"`
		WeatherCode int     `json:"weather_code"`
	} `json:"current"`
}

func GetWeather(lat, long float64) (WeatherResponse, error) {
	var weather WeatherResponse
	const weatherURL = "https://api.open-meteo.com/v1/forecast?latitude=40.71&longitude=-74.01&current=temperature_2m,relative_humidity_2m,wind_speed_10m,weather_code"
	resp, err := http.Get(weatherURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&weather)
	return weather, err
}
