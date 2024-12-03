package router

import (
	"net/http"

	_ "backend/docs" // Swagger docs import

	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler interface {
	GetTeams(w http.ResponseWriter, r *http.Request)
}

func NewRouter(handler Handler) *http.ServeMux {
	mux := http.NewServeMux()

	// Define API routes
	mux.HandleFunc("/api/teams", handler.GetTeams)

	// Serve Swagger UI
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	return mux
}
