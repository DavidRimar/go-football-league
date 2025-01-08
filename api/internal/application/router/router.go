package router

import (
	"net/http"

	_ "api/docs" // Swagger docs import

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type TeamHandler interface {
	GetTeams(w http.ResponseWriter, r *http.Request)
}

type FixtureHandler interface {
	GetFixturesByGameweek(w http.ResponseWriter, r *http.Request)
	UpdateFixture(w http.ResponseWriter, r *http.Request)
}

type TeamStatsHandler interface {
	GetStandings(w http.ResponseWriter, r *http.Request)
}

func NewRouter(teamHandler TeamHandler, fixtureHandler FixtureHandler, teamStatsHandler TeamStatsHandler) *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/api/teams", teamHandler.GetTeams).Methods(http.MethodGet)
	r.HandleFunc("/api/fixtures/{gameweekId}", fixtureHandler.GetFixturesByGameweek).Methods(http.MethodGet)
	r.HandleFunc("/api/fixtures/{fixtureId}", fixtureHandler.UpdateFixture).Methods(http.MethodPut)
	r.HandleFunc("/api/standings", teamStatsHandler.GetStandings).Methods(http.MethodGet)
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return r
}
