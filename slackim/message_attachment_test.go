package slackim

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"net/http/httptest"

	"github.com/stretchr/testify/require"
)

func TestNewAttachmentMessage(t *testing.T) {
	require.IsType(t, &AttachmentMessage{}, NewAttachmentMessage())
}

func TestAttachmentMessage_Add(t *testing.T) {
	am := &AttachmentMessage{}

	require.Empty(t, am.Attachments)
	am.Add(NewAttachment("hello"))
	require.Len(t, am.Attachments, 1)
}

func TestAttachmentMessage_Send(t *testing.T) {
	mockResponse := "mockResponse"

	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, mockResponse)
		}))
	defer ts.Close()

	b := NewAttachmentMessage().Add(NewAttachment("Hello"))

	response, err := b.Send(ts.URL)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode)

	byt, err2 := ioutil.ReadAll(response.Body)
	require.NoError(t, err2)
	require.Equal(t, mockResponse, fmt.Sprintf("%s", byt))
}
