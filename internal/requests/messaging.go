package requests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"lichess-tui/internal/errors"
	"lichess-tui/internal/requests/requestTypes"
)

func SendMessage(user string, body requestTypes.SendMessageConfig, token string) {
	url, _ := url.JoinPath("https://lichess.org/inbox", user)
	bodyBytes, _ := json.Marshal(body)

	req, err := http.NewRequest(
		POST,
		url,
		bytes.NewBuffer(bodyBytes),
	)
	if err != nil {
		errors.RequestError(err)
	}

	setHeaders(req, token, JSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.HandleRequestResponse(req, resp, err)
	}

	defer resp.Body.Close()
}
