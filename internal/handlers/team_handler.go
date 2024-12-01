package handlers

import (
	"backend/internal/models"
	"encoding/json"
	"net/http"
)

func TeamHandler(w http.ResponseWriter, r *http.Request) {
	teams := []models.Team{
		{ID: 1, Name: "Team A"},
		{ID: 2, Name: "Team B"},
		{ID: 3, Name: "Team C"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teams)
}
