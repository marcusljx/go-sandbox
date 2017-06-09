package slackim

import (
	"net/http"
)

// AttachmentMessage is the top-level message type for attachments
type AttachmentMessage struct {
	Attachments []*Attachment `json:"attachments"`
	Channel     string        `json:"channel"`
}

// NewAttachmentMessage returns a new AttachmentMessage
func NewAttachmentMessage() *AttachmentMessage {
	return &AttachmentMessage{}
}

// Add adds attachments to the AttachmentMessage
func (a *AttachmentMessage) Add(attachment ...*Attachment) *AttachmentMessage {
	a.Attachments = append(a.Attachments, attachment...)
	return a
}

// Send performs the HTTP POST request for the given message
func (a *AttachmentMessage) Send(url string) (*http.Response, error) {
	return postMessage(url, a)
}

// SendToChannel performs the HTTP POST request for the given message
func (a *AttachmentMessage) SendToChannel(url, channel string) (*http.Response, error) {
	a.Channel = channel
	return postMessage(url, a)
}
