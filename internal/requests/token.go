package requests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/proh14/lichess-tui/internal/errors"
)

type TokenInfo struct {
	Scopes  *string `json:"scopes"`
	UserID  *string `json:"userId"`
	Expires *uint64 `json:"expires"`
}

func GetTokenInfo(token string) TokenInfo {
	req := request(
		POST,
		"https://lichess.org/api/token/test",
		bytes.NewBuffer([]byte(token)),
	)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
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
