package lichess

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

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

type TokenInfo struct {
	Scopes  *string `json:"scopes"`
	UserID  *string `json:"userId"`
	Expires *int64  `json:"expires"`
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
	return GetTokenInfo(token) != TokenInfo{ nil, nil, nil }
}

// Messages
// user string
// text string
func SendMessage(body map[string]string, token string) {
	url, _ := url.JoinPath("https://lichess.org/inbox", body["user"])

	bodyBytes, _ := json.Marshal(body)
	req := request(
		POST,
		url,
		bytes.NewBuffer(bodyBytes),
	)

	setHeaders(req, token, NDJSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
	}

	defer resp.Body.Close()
}

// Game-related
// rated // bool
// time // number
// increment // number
// days // number
// variant
// ratingRange // example: 1500-1800
func SeekGame(body map[string]string, token string) {
	bodyBytes, _ := json.Marshal(body)
	req := request(POST, "https://lichess.org/api/board/seek", bytes.NewBuffer(bodyBytes))

	setHeaders(req, token, NDJSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
	}

	defer resp.Body.Close()
}

type OngoingGames struct {
	NowPlaying []struct {
		GameID   string `json:"gameId"`
		FullID   string `json:"fullId"`
		Color    string `json:"color"`
		Fen      string `json:"fen"`
		HasMoved bool   `json:"hasMoved"`
		IsMyTurn bool   `json:"isMyTurn"`
		LastMove string `json:"lastMove"`
		Opponent struct {
			ID       string `json:"id"`
			Rating   int    `json:"rating"`
			Username string `json:"username"`
		} `json:"opponent"`
		Perf        string `json:"perf"`
		Rated       bool   `json:"rated"`
		SecondsLeft int    `json:"secondsLeft"`
		Source      string `json:"source"`
		Speed       string `json:"speed"`
		Variant     struct {
			Key  string `json:"key"`
			Name string `json:"name"`
		} `json:"variant"`
	} `json:"nowPlaying"`
}

func GetOngoingGames(token string) OngoingGames {
	req := request(GET, "https://lichess.org/api/account/playing", nil)

	setHeaders(req, token, NDJSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	respMap := OngoingGames{}
	json.Unmarshal(respBody, &respMap)

	return respMap
}

const (
	OPERATION_RESIGN = "resign"
	OPERATION_ABORT  = "abort"
)

func GameOperation(gameId string, operation string, token string) {
	url, _ := url.JoinPath("https://lichess.org/api/board/game", gameId, operation)

	req := request(
		POST,
		url,
		nil,
	)

	setHeaders(req, token, NDJSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.RequestError(err)
	}

	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)

	var respMap map[string]string
	json.Unmarshal(respBody, &respMap)

	fmt.Println(respMap)
}

// Account-related
type Profile struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Perfs    struct {
		Chess960 struct {
			Games  uint  `json:"games"`
			Rating uint  `json:"rating"`
			Rd     uint  `json:"rd"`
			Prog   uint  `json:"prog"`
			Prov   bool `json:"prov"`
		} `json:"chess960"`
		Atomic struct {
			Games  uint  `json:"games"`
			Rating uint  `json:"rating"`
			Rd     uint  `json:"rd"`
			Prog   uint  `json:"prog"`
			Prov   bool `json:"prov"`
		} `json:"atomic"`
		RacingKings struct {
			Games  uint  `json:"games"`
			Rating uint  `json:"rating"`
			Rd     uint  `json:"rd"`
			Prog   uint  `json:"prog"`
			Prov   bool `json:"prov"`
		} `json:"racingKings"`
		UltraBullet struct {
			Games  uint  `json:"games"`
			Rating uint  `json:"rating"`
			Rd     uint  `json:"rd"`
			Prog   uint  `json:"prog"`
			Prov   bool `json:"prov"`
		} `json:"ultraBullet"`
		Blitz struct {
			Games  uint  `json:"games"`
			Rating uint  `json:"rating"`
			Rd     uint  `json:"rd"`
			Prog   uint  `json:"prog"`
			Prov   bool `json:"prov"`
		} `json:"blitz"`
		KingOfTheHill struct {
			Games  uint  `json:"games"`
			Rating uint  `json:"rating"`
			Rd     uint  `json:"rd"`
			Prog   uint  `json:"prog"`
			Prov   bool `json:"prov"`
		} `json:"kingOfTheHill"`
		Bullet struct {
			Games  uint  `json:"games"`
			Rating uint  `json:"rating"`
			Rd     uint  `json:"rd"`
			Prog   uint  `json:"prog"`
			Prov   bool `json:"prov"`
		} `json:"bullet"`
		Correspondence struct {
			Games  uint  `json:"games"`
			Rating uint  `json:"rating"`
			Rd     uint  `json:"rd"`
			Prog   uint  `json:"prog"`
			Prov   bool `json:"prov"`
		} `json:"correspondence"`
		Horde struct {
			Games  uint  `json:"games"`
			Rating uint  `json:"rating"`
			Rd     uint  `json:"rd"`
			Prog   uint  `json:"prog"`
			Prov   bool `json:"prov"`
		} `json:"horde"`
		Puzzle struct {
			Games  uint  `json:"games"`
			Rating uint  `json:"rating"`
			Rd     uint  `json:"rd"`
			Prog   uint  `json:"prog"`
			Prov   bool `json:"prov"`
		} `json:"puzzle"`
		Classical struct {
			Games  uint  `json:"games"`
			Rating uint  `json:"rating"`
			Rd     uint  `json:"rd"`
			Prog   uint  `json:"prog"`
			Prov   bool `json:"prov"`
		} `json:"classical"`
		Rapid struct {
			Games  uint  `json:"games"`
			Rating uint  `json:"rating"`
			Rd     uint  `json:"rd"`
			Prog   uint  `json:"prog"`
			Prov   bool `json:"prov"`
		} `json:"rapid"`
		Storm struct {
			Runs  uint `json:"runs"`
			Score uint `json:"score"`
		} `json:"storm"`
		Racer struct {
			Runs  uint `json:"runs"`
			Score uint `json:"score"`
		} `json:"racer"`
		Streak struct {
			Runs  uint `json:"runs"`
			Score uint `json:"score"`
		} `json:"streak"`
	} `json:"perfs"`
	Flair        string `json:"flair"`
	CreatedAt    uint64  `json:"createdAt"`
	Disabled     bool   `json:"disabled"`
	TosViolation bool   `json:"tosViolation"`
	Profile      struct {
		Flag       string `json:"flag"`
		Location   string `json:"location"`
		Bio        string `json:"bio"`
		RealName   string `json:"realName"`
		FideRating uint    `json:"fideRating"`
		UscfRating uint    `json:"uscfRating"`
		EcfRating  uint    `json:"ecfRating"`
		CfcRating  uint    `json:"cfcRating"`
		DsbRating  uint    `json:"dsbRating"`
		Links      string `json:"links"`
	} `json:"profile"`
	SeenAt   int64 `json:"seenAt"`
	Patron   bool  `json:"patron"`
	Verified bool  `json:"verified"`
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

	respMap := Profile{}
	json.Unmarshal(respBody, &respMap)

	return respMap
}
