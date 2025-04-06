package line

// Message represents a LINE message object
type Message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// PushRequest represents the request payload for push API
type PushRequest struct {
	To       string    `json:"to"`
	Messages []Message `json:"messages"`
}
