package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"driver-location-matching/services"
	_ "github.com/gorilla/mux"
)

type DriverController struct {
	service *services.DriverService
}

func NewDriverController(service *services.DriverService) *DriverController {
	return &DriverController{service: service}
}

func (c *DriverController) ImportDrivers(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Query().Get("filePath")
	if err := c.service.AddDriverLocations(filePath); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *DriverController) FindNearestDrivers(w http.ResponseWriter, r *http.Request) {
	lat, _ := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	lng, _ := strconv.ParseFloat(r.URL.Query().Get("lng"), 64)
	radius, _ := strconv.ParseFloat(r.URL.Query().Get("radius"), 64)

	drivers, err := c.service.FindNearestDrivers([]float64{lng, lat}, radius)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	totalNum := len(drivers)

	fmt.Println(totalNum)

	json.NewEncoder(w).Encode(drivers)
}
