package requests

import (
	"io"
	"net/http"

	"github.com/proh14/lichess-tui/internal/errors"
)

const (
	GET  = "GET"
	POST = "POST"

	NDJSON_CONTENT_TYPE = "application/x-ndjson"
	JSON_CONTENT_TYPE   = "application/json"
)

func setHeaders(req *http.Request, token string, contentType string) {
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", contentType)
}

func request(method string, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		errors.RequestError(err)
	}

	return req
}
