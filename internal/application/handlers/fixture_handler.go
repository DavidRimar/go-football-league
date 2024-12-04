package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"backend/internal/application/services"
	"backend/internal/application/utils"
)

type FixtureHandler struct {
	service *services.FixturesService
}

func NewFixtureHandler(service *services.FixturesService) *FixtureHandler {
	return &FixtureHandler{service: service}
}

func (h *FixtureHandler) validateGameweekID(id string) (int, error) {

	gameweekID, err := strconv.Atoi(id)
	if err != nil || gameweekID <= 0 {
		return 0, errors.New("invalid gameweek ID")
	}
	return gameweekID, nil
}

// @Summary Get Fixtures by Gameweek
// @Description RGet Fixtures by a Gameweek Id
// @Tags fixtures
// @Accept json
// @Produce json
// @Success 200 {array} models.Fixtures
// @Router /api/fixtures/{gameweekId} [get]
func (h *FixtureHandler) GetFixturesByGameweek(w http.ResponseWriter, r *http.Request) {

	// VALIDATE GAMEWEEK ID
	gameweekId, err := h.validateGameweekID(r.URL.Path[len("/api/fixtures/"):])
	if err != nil {
		http.Error(w, "Invalid gameweek ID", http.StatusBadRequest)
		return
	}

	// GET FIXTURES BY ID
	fixtures, err := h.service.GetFixturesByGameweek(gameweekId)
	if err != nil {
		log.Printf("Error fetching fixtures for gameweek %d: %v", gameweekId, err)
		http.Error(w, "Failed to fetch fixtures", http.StatusInternalServerError)
		return
	}

	// ENCODE RESPONSE
	utils.EncodeToJSONResponse(w, fixtures)
}
