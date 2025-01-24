package requests

import (
	"net/http"
	"net/url"

	"lichess-tui/internal/errors"
)

// https://lichess.org/api#tag/Relations/operation/followUser
func ToggleFollowUser(user string, follow bool, token string) {
	var followString string

	if follow {
		followString = "follow"
	} else {
		followString = "unfollow"
	}
	url, _ := url.JoinPath("https://lichess.org/api/rel", followString, user)

	req, err := http.NewRequest(
		POST, url, nil,
	)
	if err != nil {
		errors.RequestError(err)
	}

	setHeaders(req, token, JSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
	}
	defer resp.Body.Close()
}

// https://lichess.org/api#tag/Relations/operation/blockUser
func ToggleBlockUser(user string, block bool, token string) {
	var blockString string

	if block {
		blockString = "block"
	} else {
		blockString = "unblock"
	}
	url, _ := url.JoinPath("https://lichess.org/api/rel", blockString, user)

	req, err := http.NewRequest(
		POST, url, nil,
	)
	if err != nil {
		errors.RequestError(err)
	}

	setHeaders(req, token, JSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
	}
	defer resp.Body.Close()
}
