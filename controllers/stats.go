package controllers

import (
	"net/http"
	"meli/services"
	"encoding/json"
)

func GetStats(w http.ResponseWriter, r *http.Request)  {
	result := services.List()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}