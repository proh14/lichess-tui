package requestTypes

// https://lichess.org/api#tag/Board/operation/apiBoardSeek
type SeekGameConfig struct {
	Rated       bool   `json:"bool,omitempty"`
	Time        uint   `json:"time"`
	Increment   uint   `json:"increment"`
	Days        uint   `json:"days,omitempty"`
	Variant     string `json:"variant,omitempty"`
	RatingRange string `json:"ratingRange,omitempty"`
}

// https://lichess.org/api#tag/Games/operation/apiAccountPlaying
type OngoingGames struct {
	NowPlaying []struct {
		GameID   string `json:"gameId,omitempty"`
		FullID   string `json:"fullId,omitempty"`
		Color    string `json:"color,omitempty"`
		Fen      string `json:"fen,omitempty"`
		HasMoved bool   `json:"hasMoved,omitempty"`
		IsMyTurn bool   `json:"isMyTurn,omitempty"`
		LastMove string `json:"lastMove,omitempty"`
		Opponent struct {
			ID       string `json:"id,omitempty"`
			Rating   uint   `json:"rating,omitempty"`
			Username string `json:"username,omitempty"`
		} `json:"opponent,omitempty"`
		Perf        string `json:"perf,omitempty"`
		Rated       bool   `json:"rated,omitempty"`
		SecondsLeft uint   `json:"secondsLeft,omitempty"`
		Source      string `json:"source,omitempty"`
		Speed       string `json:"speed,omitempty"`
		Variant     struct {
			Key  string `json:"key,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"variant,omitempty"`
	} `json:"nowPlaying,omitempty"`
}

// https://lichess.org/api#tag/Board/operation/boardGameMove
type MoveConfig struct {
	OfferingDraw uint `json:"offeringDraw,omitempty"`
}

// https://lichess.org/api#tag/Board/operation/boardGameStream
type BoardState struct {
	ID      string `json:"id,omitempty"`
	Variant struct {
		Key   string `json:"key,omitempty"`
		Name  string `json:"name,omitempty"`
		Short string `json:"short,omitempty"`
	} `json:"variant,omitempty"`
	Speed string `json:"speed,omitempty"`
	Perf  struct {
		Name string `json:"name,omitempty"`
	} `json:"perf,omitempty"`
	Rated     bool   `json:"rated,omitempty"`
	CreatedAt uint64 `json:"createdAt,omitempty"`
	White     struct {
		ID     string `json:"id,omitempty"`
		Name   string `json:"name,omitempty"`
		Title  string `json:"title,omitempty"`
		Rating uint   `json:"rating,omitempty"`
	} `json:"white,omitempty"`
	Black struct {
		ID          string `json:"id,omitempty"`
		Name        string `json:"name,omitempty"`
		Title       string `json:"title,omitempty"`
		Rating      uint   `json:"rating,omitempty"`
		Provisional bool   `json:"provisional,omitempty"`
	} `json:"black,omitempty"`
	InitialFen string `json:"initialFen,omitempty"`
	Clock      struct {
		Initial   uint `json:"initial,omitempty"`
		Increment uint `json:"increment,omitempty"`
	} `json:"clock,omitempty"`
	Type  string `json:"type,omitempty"`
	State struct {
		Type   string `json:"type,omitempty"`
		Moves  string `json:"moves,omitempty"`
		Wtime  uint   `json:"wtime,omitempty"`
		Btime  uint   `json:"btime,omitempty"`
		Winc   uint   `json:"winc,omitempty"`
		Binc   uint   `json:"binc,omitempty"`
		Status string `json:"status,omitempty"`
	} `json:"state,omitempty"`
	Moves             string `json:"moves,omitempty"`
	Wtime             uint   `json:"wtime,omitempty"`
	Btime             uint   `json:"btime,omitempty"`
	Winc              uint   `json:"winc,omitempty"`
	Binc              uint   `json:"binc,omitempty"`
	Wdraw             bool   `json:"wdraw,omitempty"`
	Bdraw             bool   `json:"bdraw,omitempty"`
	Status            string `json:"status,omitempty"`
	Username          string `json:"username,omitempty"`
	Text              string `json:"text,omitempty"`
	Room              string `json:"room,omitempty"`
	Gone              bool   `json:"gone,omitempty"`
	ClaimWinInSeconds uint   `json:"claimWinInSeconds,omitempty"`
}
