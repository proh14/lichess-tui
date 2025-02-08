package requestTypes

// https://lichess.org/api#tag/Messaging/operation/inboxUsername
type SendMessageConfig struct {
	Text string `json:"text"`
}
