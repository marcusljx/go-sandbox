package slackim

import "net/http"

// BasicMessage is the top-level message type for attachments
type BasicMessage struct {
	Text string `json:"text"`
}

// NewBasicMessage returns a new BasicMessage
func NewBasicMessage(text string) *BasicMessage {
	return &BasicMessage{Text: text}
}

// Send performs the HTTP POST request for the given message
func (b *BasicMessage) Send(url string) (*http.Response, error) {
	return postMessage(url, b)
}
