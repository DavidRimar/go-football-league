package handlers

import (
	"api/internal/application/dtos"
	"api/internal/application/services"
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
// @Success 200 {array} dtos.GetTeamStatisticsDTO
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

// @Summary Update Standings
// @Description Update team statistics for specific teams
// @Tags Standings
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param updateTeamStatistics body dtos.UpdateTeamStatsDTO true "Team Statistics Update Data"
// @Success 200 {string} string "Stats updated successfully"
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Team not found"
// @Failure 500 {string} string "Failed to update statistics"
// @Router /api/standings [put]
func (h *TeamStatsHandler) UpdateStandings(w http.ResponseWriter, r *http.Request) {

	// Decode the request body
	var updateDTO dtos.UpdateTeamStatsDTO
	if err := json.NewDecoder(r.Body).Decode(&updateDTO); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Call the service to update statistics
	err := h.service.UpdateTeamStatistics(r.Context(), updateDTO)
	if err != nil {
		http.Error(w, "Failed to update statistics", http.StatusInternalServerError)
		return
	}

	// Send a success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Stats updated successfully"}`))
}
