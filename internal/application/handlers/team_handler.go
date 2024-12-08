package handlers

import (
	"backend/internal/application/services"
	"encoding/json"
	"net/http"
)

type TeamHandler struct {
	service *services.TeamService
}

func NewTeamHandler(service *services.TeamService) *TeamHandler {
	return &TeamHandler{service: service}
}

// @Summary Get all teams
// @Description Retrieve details of all teams
// @Tags Teams
// @Accept json
// @Produce json
// @Success 200 {array} models.Team
// @Router /api/teams [get]
func (h *TeamHandler) GetTeams(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	teams, err := h.service.GetTeams(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch teams", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teams)
}
