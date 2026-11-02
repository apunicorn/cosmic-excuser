package main

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

const maxActionLength = 500

type excuseRequest struct {
	Action string `json:"action"`
	Sign   string `json:"sign"`
}

var zodiacSigns = []string{
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
		http.Error(w, "use POST", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 4096)
	var req excuseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	req.Action = strings.TrimSpace(req.Action)
	if req.Action == "" || len(req.Action) > maxActionLength {
		http.Error(w, "action must be 1-500 characters", http.StatusBadRequest)
		return
	}

	if req.Sign == "" {
		http.Error(w, "No sign chosen", http.StatusBadRequest)
		return
	}

	if !slices.Contains(zodiacSigns, req.Sign) {
		http.Error(w, "Invalid sign chosen", http.StatusBadRequest)
		return
	}

	excuse, err := GenerateExcuse(req.Action, req.Sign)
	if err != nil {
		log.Println("generate excuse:", err)
		http.Error(w, "could not generate an excuse, try again", http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"excuse": excuse})
}

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
	http.HandleFunc("/api/excuse", ExcuseAPI)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
