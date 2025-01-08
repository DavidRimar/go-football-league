package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"api/internal/application/dtos"
	"api/internal/application/services"
	"api/internal/application/utils"
	"api/internal/domain/interfaces"
	"api/internal/domain/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type FixtureHandler struct {
	fixtureService   *services.FixturesService
	teamStatsService *services.TeamStatsService
	eventPublisher   interfaces.EventPublisher
}

func NewFixtureHandler(fixtureService *services.FixturesService,
	teamStatsService *services.TeamStatsService,
	eventPublisher interfaces.EventPublisher) *FixtureHandler {
	return &FixtureHandler{
		fixtureService:   fixtureService,
		teamStatsService: teamStatsService,
		eventPublisher:   eventPublisher,
	}
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
	fixtures, err := h.fixtureService.GetFixturesByGameweek(r.Context(), gameweekId)
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

	// Fetch the existing fixture
	fixture, err := h.fixtureService.GetFixtureByID(ctx, fixtureID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			http.Error(w, "Fixture not found", http.StatusNotFound)
		}
	}

	// Check if the fixture's status is "Played"
	if fixture.Status == models.StatusPlayed {
		http.Error(w, "Cannot update a fixture that is already marked as Played", http.StatusBadRequest)
		return
	}

	// Call the service to update the fixture
	if err := h.fixtureService.UpdateFixture(ctx, fixtureID, dto); err != nil {
		http.Error(w, "Failed to update fixture", http.StatusInternalServerError)
		return
	}

	// PUBLISH EVENT
	if dto.Status == models.StatusPlayed {

		newFixture := dtos.UpdateTeamStatsDTO{
			HomeTeamId: fixture.HomeTeamId,
			HomeScore:  dto.HomeScore,
			AwayTeamId: fixture.AwayTeamId,
			AwayScore:  dto.AwayScore,
		}

		err := h.eventPublisher.PublishEvent(newFixture)
		if err != nil {
			log.Printf("Failed to publish event: %v", err)
		}
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Fixture updated successfully"))
}
