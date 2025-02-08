package requestTypes

// https://lichess.org/api#tag/OAuth/operation/tokenTest
type TokenInfo struct {
	// Using pointers in order to handle null
	Scopes  *string `json:"scopes"`
	UserID  *string `json:"userId"`
	Expires *uint64 `json:"expires"`
}
