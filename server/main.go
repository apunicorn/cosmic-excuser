package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func WeatherAPI(w http.ResponseWriter, r *http.Request) {
	weather, err := GetWeather(lat, long)
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
