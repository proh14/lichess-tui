package requests

import (
	"encoding/json"
	"net/http"
	"lichess-tui/internal/errors"
)

var IncomingEvents map[string]string

func StreamIncomingEvents(streamVar *map[string]string, token string) {
	req := request(GET, "https://lichess.org/api/stream/event", nil)

	setHeaders(req, token, NDJSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	
	for {
		dec.Decode(&streamVar)
	}
}

