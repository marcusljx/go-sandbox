package slackim

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"net/http/httptest"

	"github.com/stretchr/testify/require"
)

func TestBasicMessage_Send(t *testing.T) {
	mockRequest := "Hello World"
	mockResponse := "mockResponse"

	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, mockResponse)
		}))
	defer ts.Close()

	b := NewBasicMessage(mockRequest)
	require.Equal(t, mockRequest, b.Text)

	response, err := b.Send(ts.URL)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode)

	byt, err2 := ioutil.ReadAll(response.Body)
	require.NoError(t, err2)
	require.Equal(t, mockResponse, fmt.Sprintf("%s", byt))
}
