package lichess

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"github.com/proh14/lichess-tui/internal/errors"
)

func TokenExists(token string) bool {
	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	url := "https://lichess.org/api/account"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errors.RequestError(err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	var respMap map[string]string
	json.Unmarshal(respBody, &respMap)

	// An error is returned in case a token doesn't exist
	_, containsKey := respMap["error"]

	return !containsKey
}

func SendMessage(user string, text string, token string) {
	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	body := map[string]string{
		"text": text,
	}

	bodyBytes, _ := json.Marshal(body)

	url := "https://lichess.org/inbox/" + user
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		errors.RequestError(err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}

	defer resp.Body.Close()
}
