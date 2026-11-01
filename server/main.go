package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func WeatherAPI(w http.ResponseWriter, r *http.Request) {
	lat, latErr := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	lon, longErr := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)
	if latErr != nil || longErr != nil {
		http.Error(w, "lat and lon are required numbers", http.StatusBadRequest)
		return
	}
	weather, err := GetWeather(lat, lon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)

}

func main() {
	http.HandleFunc("/api/weather", WeatherAPI)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
