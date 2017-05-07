package slackim

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Monospace returns text enclosed with the monospace marker('):
//   `text`
func Monospace(text interface{}) string {
	return fmt.Sprintf("`%v`", text)
}

// Code returns the text enclosed with the code marker (```)
//   ```
//   text
//   ```
func Code(text interface{}) string {
	return fmt.Sprintf("```\n%v\n```", text)
}

// Hyperlink returns the string for encoding the text with the given url hyperlink
func Hyperlink(text, url string) string {
	return fmt.Sprintf("<%s|%s>", url, text)
}

//============================================================ UnEXPORTED

// findInList searches a list for an item, and returns the position of the item if found.
// If the item does not exist in the list, this function returns -1
func findInList(list []string, item string) int {
	var found bool
	var i int
	var m string
	for i, m = range list {
		if m == item {
			found = true
			break
		}
	}
	if !found {
		return -1
	}
	return i
}

// postMessage POSTs a payload to the URL
func postMessage(url string, msg interface{}) (*http.Response, error) {
	payload := new(bytes.Buffer)
	err := json.NewEncoder(payload).Encode(msg)
	if err != nil {
		log.Printf("error while encoding msg : %v", err)
		return nil, err
	}

	log.Printf("sending slackim notification: %v", payload)
	return http.Post(url, "application/json", payload)
}
