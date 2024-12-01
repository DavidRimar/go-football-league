package main

import (
	"backend/internal/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/teams", handlers.TeamHandler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
