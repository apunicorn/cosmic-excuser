package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// these all get rejected before any outside API is called
func TestExcuseAPIRejectsBadRequests(t *testing.T) {
	tests := []struct {
		name   string
		method string
		body   string
		want   int
	}{
		{"wrong method", http.MethodGet, ``, http.StatusMethodNotAllowed},
		{"bad json", http.MethodPost, `{not json`, http.StatusBadRequest},
		{"empty action", http.MethodPost, `{"action":"  ","sign":"Leo"}`, http.StatusBadRequest},
		{"action too long", http.MethodPost, `{"action":"` + strings.Repeat("a", maxActionLength+1) + `","sign":"Leo"}`, http.StatusBadRequest},
		{"missing sign", http.MethodPost, `{"action":"ate coworker leftovers"}`, http.StatusBadRequest},
		{"fake sign", http.MethodPost, `{"action":"ate coworker leftovers","sign":"Dragon"}`, http.StatusBadRequest},
		{"lat out of range", http.MethodPost, `{"action":"ate coworker leftovers","sign":"Leo","lat":999,"lon":0}`, http.StatusBadRequest},
		{"lon out of range", http.MethodPost, `{"action":"ate coworker leftovers","sign":"Leo","lat":0,"lon":-181}`, http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/api/excuse", strings.NewReader(tt.body))
			rec := httptest.NewRecorder()

			ExcuseAPI(rec, req)

			if rec.Code != tt.want {
				t.Errorf("got status %d, want %d (body: %q)", rec.Code, tt.want, rec.Body.String())
			}
		})
	}
}

func TestWeatherAPIRequiresCoordinates(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/weather?lat=abc", nil)
	rec := httptest.NewRecorder()

	WeatherAPI(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("got status %d, want %d", rec.Code, http.StatusBadRequest)
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		c, want float64
	}{
		{0, 32},
		{100, 212},
		{-40, -40}, // F matches C at -40
	}

	for _, tt := range tests {
		if got := celsiusToFahrenheit(tt.c); got != tt.want {
			t.Errorf("celsiusToFahrenheit(%v) = %v, want %v", tt.c, got, tt.want)
		}
	}
}

func TestGetWeatherDetailsFromAPI(t *testing.T) {
	var w WeatherResponse
	w.Current.Temperature = 0
	w.Current.Humidity = 61
	w.Current.WindSpeed = 9.4
	w.Current.WeatherCode = 3

	got := getWeatherDetailsFromAPI(w)
	want := "32.0°F, 61% humidity, wind 9 km/h, weather code 3"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGenerateExcuseNeedsAPIKey(t *testing.T) {
	t.Setenv("AI_API_KEY", "") // make sure we never hit the real API

	if _, err := GenerateExcuse("ate coworker leftovers", "Leo", nil); err == nil {
		t.Error("expected an error when AI_API_KEY is not set")
	}
}
