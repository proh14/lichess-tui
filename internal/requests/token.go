package requests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"lichess-tui/internal/errors"
)

// https://lichess.org/api#tag/OAuth/operation/tokenTest
type TokenInfo struct {
	// Using pointers in order to handle null
	Scopes  *string `json:"scopes"`
	UserID  *string `json:"userId"`
	Expires *uint64 `json:"expires"`
}

func GetTokenInfo(token string) TokenInfo {
	req, err := http.NewRequest(
		POST,
		"https://lichess.org/api/token/test",
		bytes.NewBuffer([]byte(token)),
	)
	if err != nil {
		errors.RequestError(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.HandleRequestResponse(req, resp, err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	var respMap map[string]TokenInfo
	json.Unmarshal(respBody, &respMap)

	return respMap[token]
}

func TokenExists(token string) bool {
	return GetTokenInfo(token) != TokenInfo{nil, nil, nil}
}
