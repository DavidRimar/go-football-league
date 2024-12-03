package utils

import (
	"backend/internal/domain/models"
	"encoding/json"
	"os"
)

func LoadTeamsFromJSON(filePath string) []models.Team {

	byteValue, err := os.ReadFile(filePath)

	if err != nil {
		return nil
	}

	var teams []models.Team
	if err := json.Unmarshal(byteValue, &teams); err != nil {
		return nil
	}

	if len(teams) == 0 {
		return nil
	}

	return teams
}
