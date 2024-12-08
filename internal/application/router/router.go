package router

import (
	"net/http"

	_ "backend/docs" // Swagger docs import

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

func NewRouter(teamHandler TeamHandler, fixtureHandler FixtureHandler) *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/api/teams", teamHandler.GetTeams).Methods("GET")
	r.HandleFunc("/api/fixtures/{gameweekId}", fixtureHandler.GetFixturesByGameweek).Methods("GET")
	r.HandleFunc("/api/fixture/{fixtureId}", fixtureHandler.UpdateFixture).Methods("PUT")
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return r
}
