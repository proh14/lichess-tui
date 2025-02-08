package requests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"lichess-tui/internal/errors"
	"lichess-tui/internal/requests/requestTypes"
)

func GetTokenInfo(token string) requestTypes.TokenInfo {
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

	var respMap map[string]requestTypes.TokenInfo
	json.Unmarshal(respBody, &respMap)

	return respMap[token]
}

func TokenExists(token string) bool {
	return GetTokenInfo(token) != requestTypes.TokenInfo{nil, nil, nil}
}
