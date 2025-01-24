package requests

import (
	"net/http"
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
