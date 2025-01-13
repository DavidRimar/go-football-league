package utils

import (
	"bytes"
	"fmt"
	"net/http"
)

func CallAPI(method, url string, body []byte) error {

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create %s request: %w", method, err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Use the HTTP client to send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make API request: %w", err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	return nil
}
