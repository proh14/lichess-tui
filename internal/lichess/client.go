package lichess

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

func ValidateToken(value string) error {
	if len(value) == 0 {
		return errors.New("The token can't be empty.")
	}

	headers := map[string]string{
		"Authorization": "Bearer " + value,
		"Content-Type":  "application/json",
	}

	// Make the POST request
	url := "https://lichess.org/api/account"
	req, _ := http.NewRequest("GET", url, nil)

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	var respMap map[string]string
	json.Unmarshal(respBody, &respMap)

	value, containsKey := respMap["error"]

	if containsKey {
		return errors.New("The given token is invalid.")
	} else {
		return nil
	}
}
