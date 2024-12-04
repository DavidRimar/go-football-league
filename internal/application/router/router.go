package router

import (
	"net/http"

	_ "backend/docs" // Swagger docs import

	httpSwagger "github.com/swaggo/http-swagger"
)

type TeamHandler interface {
	GetTeams(w http.ResponseWriter, r *http.Request)
}

type FixtureHandler interface {
	GetFixturesByGameweek(w http.ResponseWriter, r *http.Request)
}

func NewRouter(teamHandler TeamHandler, fixtureHandler FixtureHandler) *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/api/teams", teamHandler.GetTeams)
	mux.HandleFunc("/api/fixtures/", fixtureHandler.GetFixturesByGameweek)
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	return mux
}
