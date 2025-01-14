package lichess

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/proh14/lichess-tui/internal/errors"
)

const (
	GET  = "GET"
	POST = "POST"
)

func setHeaders(req *http.Request, token string) {
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
}

func request(method string, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		errors.RequestError(err)
	}

	return req
}

func TokenExists(token string) bool {
	req := request(GET, "https://lichess.org/api/account", nil)

	setHeaders(req, token)

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

// Messages
type SendMessageConfig struct {
	user string
	text string
}

func SendMessage(config SendMessageConfig, token string) {
	body := map[string]string{
		"text": config.text,
	}
	bodyBytes, _ := json.Marshal(body)
	req := request(POST, "https://lichess.org/inbox/"+config.user, bytes.NewBuffer(bodyBytes))

	setHeaders(req, token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
	}

	defer resp.Body.Close()
}

type SeekGameConfig struct {
	rated       string // bool
	time        string // number
	increment   string // number
	days        string // number
	variant     string
	ratingRange string // example: 1500-1800
}

// Game operations
func SeekGame(config SeekGameConfig, token string) {
	body := map[string]string{
		"rated":       config.rated,
		"time":        config.time,
		"increment":   config.increment,
		"days":        config.days,
		"variant":     config.variant,
		"ratingRange": config.ratingRange,
	}
	bodyBytes, _ := json.Marshal(body)
	req := request(POST, "https://lichess.org/api/board/seek", bytes.NewBuffer(bodyBytes))

	setHeaders(req, token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
	}

	defer resp.Body.Close()
}
