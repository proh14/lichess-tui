package requests

import (
	"encoding/json"
	"io"
	"net/http"

	"lichess-tui/internal/errors"
	"lichess-tui/internal/requests/requestTypes"
)

func GetProfile(token string) requestTypes.Profile {
	req, err := http.NewRequest(
		GET, "https://lichess.org/api/account", nil,
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

	respBody, _ := io.ReadAll(resp.Body)

	var respMap requestTypes.Profile
	json.Unmarshal(respBody, &respMap)

	return respMap
}

func GetEmailAddress(token string) requestTypes.EmailAddress {
	req, err := http.NewRequest(
		GET, "https://lichess.org/api/account/email", nil,
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

	respBody, _ := io.ReadAll(resp.Body)

	var respMap requestTypes.EmailAddress
	json.Unmarshal(respBody, &respMap)

	return respMap
}
