package requests

import (
	"encoding/json"
	"io"
	"net/http"

	"lichess-tui/internal/errors"
)

// https://lichess.org/api#tag/Account/operation/accountMe
type Profile struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Perfs    struct {
		Chess960 struct {
			Games  uint `json:"games,omitempty"`
			Rating uint `json:"rating,omitempty"`
			Rd     uint `json:"rd,omitempty"`
			Prog   uint `json:"prog,omitempty"`
			Prov   bool `json:"prov,omitempty"`
		} `json:"chess960,omitempty"`
		Atomic struct {
			Games  uint `json:"games,omitempty"`
			Rating uint `json:"rating,omitempty"`
			Rd     uint `json:"rd,omitempty"`
			Prog   uint `json:"prog,omitempty"`
			Prov   bool `json:"prov,omitempty"`
		} `json:"atomic,omitempty"`
		RacingKings struct {
			Games  uint `json:"games,omitempty"`
			Rating uint `json:"rating,omitempty"`
			Rd     uint `json:"rd,omitempty"`
			Prog   uint `json:"prog,omitempty"`
			Prov   bool `json:"prov,omitempty"`
		} `json:"racingKings,omitempty"`
		UltraBullet struct {
			Games  uint `json:"games,omitempty"`
			Rating uint `json:"rating,omitempty"`
			Rd     uint `json:"rd,omitempty"`
			Prog   uint `json:"prog,omitempty"`
			Prov   bool `json:"prov,omitempty"`
		} `json:"ultraBullet,omitempty"`
		Blitz struct {
			Games  uint `json:"games,omitempty"`
			Rating uint `json:"rating,omitempty"`
			Rd     uint `json:"rd,omitempty"`
			Prog   uint `json:"prog,omitempty"`
			Prov   bool `json:"prov,omitempty"`
		} `json:"blitz,omitempty"`
		KingOfTheHill struct {
			Games  uint `json:"games,omitempty"`
			Rating uint `json:"rating,omitempty"`
			Rd     uint `json:"rd,omitempty"`
			Prog   uint `json:"prog,omitempty"`
			Prov   bool `json:"prov,omitempty"`
		} `json:"kingOfTheHill,omitempty"`
		Bullet struct {
			Games  uint `json:"games,omitempty"`
			Rating uint `json:"rating,omitempty"`
			Rd     uint `json:"rd,omitempty"`
			Prog   uint `json:"prog,omitempty"`
			Prov   bool `json:"prov,omitempty"`
		} `json:"bullet,omitempty"`
		Correspondence struct {
			Games  uint `json:"games,omitempty"`
			Rating uint `json:"rating,omitempty"`
			Rd     uint `json:"rd,omitempty"`
			Prog   uint `json:"prog,omitempty"`
			Prov   bool `json:"prov,omitempty"`
		} `json:"correspondence,omitempty"`
		Horde struct {
			Games  uint `json:"games,omitempty"`
			Rating uint `json:"rating,omitempty"`
			Rd     uint `json:"rd,omitempty"`
			Prog   uint `json:"prog,omitempty"`
			Prov   bool `json:"prov,omitempty"`
		} `json:"horde,omitempty"`
		Puzzle struct {
			Games  uint `json:"games,omitempty"`
			Rating uint `json:"rating,omitempty"`
			Rd     uint `json:"rd,omitempty"`
			Prog   uint `json:"prog,omitempty"`
			Prov   bool `json:"prov,omitempty"`
		} `json:"puzzle,omitempty"`
		Classical struct {
			Games  uint `json:"games,omitempty"`
			Rating uint `json:"rating,omitempty"`
			Rd     uint `json:"rd,omitempty"`
			Prog   uint `json:"prog,omitempty"`
			Prov   bool `json:"prov,omitempty"`
		} `json:"classical,omitempty"`
		Rapid struct {
			Games  uint `json:"games,omitempty"`
			Rating uint `json:"rating,omitempty"`
			Rd     uint `json:"rd,omitempty"`
			Prog   uint `json:"prog,omitempty"`
			Prov   bool `json:"prov,omitempty"`
		} `json:"rapid,omitempty"`
		Storm struct {
			Runs  uint `json:"runs,omitempty"`
			Score uint `json:"score,omitempty"`
		} `json:"storm,omitempty"`
		Racer struct {
			Runs  uint `json:"runs,omitempty"`
			Score uint `json:"score,omitempty"`
		} `json:"racer,omitempty"`
		Streak struct {
			Runs  uint `json:"runs,omitempty"`
			Score uint `json:"score,omitempty"`
		} `json:"streak,omitempty"`
	} `json:"perfs,omitempty"`
	Flair        string `json:"flair,omitempty"`
	CreatedAt    uint64 `json:"createdAt,omitempty"`
	Disabled     bool   `json:"disabled,omitempty"`
	TosViolation bool   `json:"tosViolation,omitempty"`
	Profile      struct {
		Flag       string `json:"flag,omitempty"`
		Location   string `json:"location,omitempty"`
		Bio        string `json:"bio,omitempty"`
		RealName   string `json:"realName,omitempty"`
		FideRating uint   `json:"fideRating,omitempty"`
		UscfRating uint   `json:"uscfRating,omitempty"`
		EcfRating  uint   `json:"ecfRating,omitempty"`
		CfcRating  uint   `json:"cfcRating,omitempty"`
		DsbRating  uint   `json:"dsbRating,omitempty"`
		Links      string `json:"links,omitempty"`
	} `json:"profile,omitempty"`
	SeenAt   uint64 `json:"seenAt,omitempty"`
	Patron   bool   `json:"patron,omitempty"`
	Verified bool   `json:"verified,omitempty"`
	PlayTime struct {
		Total uint `json:"total,omitempty"`
		Tv    uint `json:"tv,omitempty"`
	} `json:"playTime,omitempty"`
	Title   string `json:"title,omitempty"`
	URL     string `json:"url,omitempty"`
	Playing string `json:"playing,omitempty"`
	Count   struct {
		All      uint `json:"all,omitempty"`
		Rated    uint `json:"rated,omitempty"`
		Ai       uint `json:"ai,omitempty"`
		Draw     uint `json:"draw,omitempty"`
		DrawH    uint `json:"drawH,omitempty"`
		Loss     uint `json:"loss,omitempty"`
		LossH    uint `json:"lossH,omitempty"`
		Win      uint `json:"win,omitempty"`
		WinH     uint `json:"winH,omitempty"`
		Bookmark uint `json:"bookmark,omitempty"`
		Playing  uint `json:"playing,omitempty"`
		Import   uint `json:"import,omitempty"`
		Me       uint `json:"me,omitempty"`
	} `json:"count,omitempty"`
	Streaming bool `json:"streaming,omitempty"`
	Streamer  struct {
		Twitch struct {
			Channel string `json:"channel,omitempty"`
		} `json:"twitch,omitempty"`
		YouTube struct {
			Channel string `json:"channel,omitempty"`
		} `json:"youTube,omitempty"`
	} `json:"streamer,omitempty"`
	Followable bool `json:"followable,omitempty"`
	Following  bool `json:"following,omitempty"`
	Blocking   bool `json:"blocking,omitempty"`
	FollowsYou bool `json:"followsYou,omitempty"`
}

func GetProfile(token string) Profile {
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
		errors.RequestError(err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	var respMap Profile
	json.Unmarshal(respBody, &respMap)

	return respMap
}

type EmailAddress struct {
	Email string `json:"email"`
}

func GetEmailAddress(token string) EmailAddress {
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
		errors.RequestError(err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	var respMap EmailAddress
	json.Unmarshal(respBody, &respMap)

	return respMap
}
