package slackim

import "net/http"

// Message is the standard message interface
type Message interface {
	Send(url string) (*http.Response, error)
}
