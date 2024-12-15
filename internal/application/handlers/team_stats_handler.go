package handlers

import (
	"backend/internal/application/services"
	"encoding/json"
	"net/http"
)

type TeamStatsHandler struct {
	service *services.TeamStatsService
}

func NewTeamStatsHandler(service *services.TeamStatsService) *TeamStatsHandler {
	return &TeamStatsHandler{service: service}
}

// @Summary Get standings
// @Description Retrieve team statistics for all teams
// @Tags Standings
// @Accept json
// @Produce json
// @Success 200 {array} models.TeamStatistics
// @Router /api/standings [get]
func (h *TeamStatsHandler) GetStandings(w http.ResponseWriter, r *http.Request) {

	stats, err := h.service.GetTeamStatistics(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch stats", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
