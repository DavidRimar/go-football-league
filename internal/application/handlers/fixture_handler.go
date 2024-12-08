package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"backend/internal/application/dtos"
	"backend/internal/application/services"
	"backend/internal/application/utils"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
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
// @Description Get Fixtures by a Gameweek Id
// @Tags Fixtures
// @Accept json
// @Produce json
// @Param gameweekId path string true "Gameweek"
// @Success 200 {array} models.Fixture
// @Router /api/fixtures/{gameweekId} [get]
func (h *FixtureHandler) GetFixturesByGameweek(w http.ResponseWriter, r *http.Request) {

	// VALIDATE GAMEWEEK ID
	gameweekId, err := h.validateGameweekID(r.URL.Path[len("/api/fixtures/"):])
	if err != nil {
		http.Error(w, "Invalid gameweek ID", http.StatusBadRequest)
		return
	}

	// GET FIXTURES BY ID
	fixtures, err := h.service.GetFixturesByGameweek(r.Context(), gameweekId)
	if err != nil {
		log.Printf("Error fetching fixtures for gameweek %d: %v", gameweekId, err)
		http.Error(w, "Failed to fetch fixtures", http.StatusInternalServerError)
		return
	}

	// ENCODE RESPONSE
	utils.EncodeToJSONResponse(w, fixtures)
}

// @Summary Update Fixture
// @Description Update the fixture's status and scores by its ID.
// @Tags Fixtures
// @Accept json
// @Produce json
// @Param fixtureId path string true "Fixture ID"
// @Param body body dtos.UpdateFixtureDTO true "Fixture Update Payload"
// @Success 200 {string} string "Fixture updated successfully"
// @Failure 400 {string} string "Invalid request body or missing fixture ID"
// @Failure 404 {string} string "Fixture not found"
// @Router /api/fixtures/{fixtureId} [put]
func (h *FixtureHandler) UpdateFixture(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// Extract fixture ID from the URL
	vars := mux.Vars(r)
	fixtureID, ok := vars["fixtureId"]
	if !ok || fixtureID == "" {
		http.Error(w, "Fixture ID is required", http.StatusBadRequest)
		return
	}

	// Parse request body
	var dto dtos.UpdateFixtureDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call the service to update the fixture
	if err := h.service.UpdateFixture(ctx, fixtureID, dto); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			http.Error(w, "Fixture not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to update fixture", http.StatusInternalServerError)
		}
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Fixture updated successfully"))
}
