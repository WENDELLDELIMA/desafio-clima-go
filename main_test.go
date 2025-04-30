package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestWeatherHandler_InvalidCEP(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/weather/123", nil)
	w := httptest.NewRecorder()
	WeatherHandler(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("Expected status 422, got %d", w.Code)
	}
}

func TestWeatherHandler_NotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/weather/00000000", nil)
	w := httptest.NewRecorder()
	WeatherHandler(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

func TestMain(m *testing.M) {
	os.Setenv("WEATHER_API_KEY", "test_key")
	os.Exit(m.Run())
}
