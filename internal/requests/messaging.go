package requests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/proh14/lichess-tui/internal/errors"
)

// https://lichess.org/api#tag/Messaging/operation/inboxUsername
type SendMessageConfig struct {
	Text string `json:"text,omitempty"`
}

func SendMessage(user string, body SendMessageConfig, token string) {
	url, _ := url.JoinPath("https://lichess.org/inbox", user)

	bodyBytes, _ := json.Marshal(body)
	req := request(
		POST,
		url,
		bytes.NewBuffer(bodyBytes),
	)

	setHeaders(req, token, NDJSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
	}

	defer resp.Body.Close()
}
