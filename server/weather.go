package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type WeatherResponse struct {
	Current struct {
		Temperature float64 `json:"temperature_2m"`
		Humidity    int     `json:"relative_humidity_2m"`
		WindSpeed   float64 `json:"wind_speed_10m"`
		WeatherCode int     `json:"weather_code"`
	} `json:"current"`
}

func GetWeather(lat, lon float64) (WeatherResponse, error) {
	var weather WeatherResponse

	query := url.Values{}
	query.Set("latitude", strconv.FormatFloat(lat, 'f', 4, 64))
	query.Set("longitude", strconv.FormatFloat(lon, 'f', 4, 64))

	query.Set("current", "temperature_2m,relative_humidity_2m,wind_speed_10m,weather_code")
	weatherURL := "https://api.open-meteo.com/v1/forecast?" + query.Encode()

	resp, err := http.Get(weatherURL)
	if err != nil {
		return weather, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return weather, fmt.Errorf("open-meteo returned %s", resp.Status)
	}
	err = json.NewDecoder(resp.Body).Decode(&weather)
	return weather, err
}
