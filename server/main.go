package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

const maxActionLength = 500 // cap on what the user can type. don't want a thesis

type excuseRequest struct {
	Action string   `json:"action"`
	Sign   string   `json:"sign"`
	Lat    *float64 `json:"lat"` // pointers so location is optional
	Lon    *float64 `json:"lon"`
}

var zodiacSigns = []string{ // zodiac signs are not up for question
	"Aries",
	"Taurus",
	"Gemini",
	"Cancer",
	"Leo",
	"Virgo",
	"Libra",
	"Scorpio",
	"Sagittarius",
	"Capricorn",
	"Aquarius",
	"Pisces",
}

func ExcuseAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "use POST", http.StatusMethodNotAllowed) // 405 wrong method
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 4096) // limit request size
	var req excuseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		var tooBig *http.MaxBytesError
		if errors.As(err, &tooBig) {
			http.Error(w, "request too large", http.StatusRequestEntityTooLarge) // 413 err
			return
		}
		http.Error(w, "invalid JSON", http.StatusBadRequest) // 400 bad request
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest) // 400 bad request
		return
	}
	req.Action = strings.TrimSpace(req.Action)
	if req.Action == "" || len(req.Action) > maxActionLength {
		http.Error(w, "action must be 1-500 characters", http.StatusBadRequest) // 400 bad request
		return
	}

	if req.Sign == "" {
		http.Error(w, "No sign chosen", http.StatusBadRequest) // 400 bad request
		return
	}

	if !slices.Contains(zodiacSigns, req.Sign) {
		http.Error(w, "Invalid sign chosen", http.StatusBadRequest) // 400 bad request
		return
	}

	var weather *WeatherResponse // stays nil if no location or lookup fails
	if req.Lat != nil && req.Lon != nil {
		lat, lon := *req.Lat, *req.Lon
		if (lat > 90 || lat < -90) || (lon > 180 || lon < -180) { // valid coords range
			http.Error(w, "lat/long out of bounds", http.StatusBadRequest)
			return
		}
		wthr, err := GetWeather(*req.Lat, *req.Lon)
		if err != nil {
			log.Println("weather error: ", err) // carry on without weather
		} else {
			weather = &wthr
		}
	}

	excuse, err := GenerateExcuse(req.Action, req.Sign, weather)
	if err != nil {
		log.Println("generate excuse:", err)
		http.Error(w, "could not generate an excuse, try again", http.StatusBadGateway) // 502
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"excuse": excuse})
}

func WeatherAPI(w http.ResponseWriter, r *http.Request) { // get the weather using lat/long from user geo
	lat, latErr := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	lon, longErr := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)

	if latErr != nil || longErr != nil {
		http.Error(w, "lat and lon are required numbers", http.StatusBadRequest) // 400 bad request
		return
	}

	weather, err := GetWeather(lat, lon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway) // 502
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)

}

func main() {
	http.HandleFunc("/api/weather", WeatherAPI)
	http.HandleFunc("/api/excuse", ExcuseAPI)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
