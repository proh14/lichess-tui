package requests

import (
	"encoding/json"
	"io"
	"net/http"

	"lichess-tui/internal/errors"
)

// https://lichess.org/api#tag/Account/operation/accountMe
type Profile struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Perfs    struct {
		Chess960, Atomic, RacingKings, UltraBullet, Blitz, KingOfTheHill, Bullet, Correspondence, Horde, Puzzle, Classical, Rapid, Storm, Racer, Streak struct {
			Games  uint `json:"games"`
			Rating uint `json:"rating"`
			Rd     uint `json:"rd"`
			Prog   uint `json:"prog"`
			Prov   bool `json:"prov"`
		} `json:"chess960"`
	} `json:"perfs"`
	Flair        string `json:"flair"`
	CreatedAt    uint64 `json:"createdAt"`
	Disabled     bool   `json:"disabled"`
	TosViolation bool   `json:"tosViolation"`
	Profile      struct {
		Flag       string `json:"flag"`
		Location   string `json:"location"`
		Bio        string `json:"bio"`
		RealName   string `json:"realName"`
		FideRating uint   `json:"fideRating"`
		UscfRating uint   `json:"uscfRating"`
		EcfRating  uint   `json:"ecfRating"`
		CfcRating  uint   `json:"cfcRating"`
		DsbRating  uint   `json:"dsbRating"`
		Links      string `json:"links"`
	} `json:"profile"`
	SeenAt   uint64 `json:"seenAt"`
	Patron   bool   `json:"patron"`
	Verified bool   `json:"verified"`
	PlayTime struct {
		Total uint `json:"total"`
		Tv    uint `json:"tv"`
	} `json:"playTime"`
	Title   string `json:"title"`
	URL     string `json:"url"`
	Playing string `json:"playing"`
	Count   struct {
		All      uint `json:"all"`
		Rated    uint `json:"rated"`
		Ai       uint `json:"ai"`
		Draw     uint `json:"draw"`
		DrawH    uint `json:"drawH"`
		Loss     uint `json:"loss"`
		LossH    uint `json:"lossH"`
		Win      uint `json:"win"`
		WinH     uint `json:"winH"`
		Bookmark uint `json:"bookmark"`
		Playing  uint `json:"playing"`
		Import   uint `json:"import"`
		Me       uint `json:"me"`
	} `json:"count"`
	Streaming bool `json:"streaming"`
	Streamer  struct {
		Twitch struct {
			Channel string `json:"channel"`
		} `json:"twitch"`
		YouTube struct {
			Channel string `json:"channel"`
		} `json:"youTube"`
	} `json:"streamer"`
	Followable bool `json:"followable"`
	Following  bool `json:"following"`
	Blocking   bool `json:"blocking"`
	FollowsYou bool `json:"followsYou"`
}

func GetProfile(token string) Profile {
	req := request(GET, "https://lichess.org/api/account", nil)

	setHeaders(req, token, NDJSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	var respMap Profile
	json.Unmarshal(respBody, &respMap)

	return respMap
}
