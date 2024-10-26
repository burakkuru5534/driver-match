package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"driver-location-match/services"
)

type MatchingController struct {
	service *services.MatchingService
}

func NewMatchingController(service *services.MatchingService) *MatchingController {
	return &MatchingController{service: service}
}

func (c *MatchingController) FindNearestDrivers(w http.ResponseWriter, r *http.Request) {
	lat, _ := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	lng, _ := strconv.ParseFloat(r.URL.Query().Get("lng"), 64)
	radius, _ := strconv.ParseFloat(r.URL.Query().Get("radius"), 64)

	drivers, err := c.service.FindNearestDrivers([]float64{lng, lat}, radius)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(drivers)
}
