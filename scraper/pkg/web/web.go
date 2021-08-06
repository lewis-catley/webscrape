package web

import (
	"io"
	"net/http"
	"net/url"
)

// HTTPClient custom abstraction of http to allow us to mock the implementation
type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

// Client stores the instance of HTTPClient, exposed for ease of testing
var Client HTTPClient

func init() {
	Client = &http.Client{}
}

// GetURLContent will get the page content from a given url
func GetURLContent(u url.URL) (*[]byte, error) {
	resp, err := Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return &body, err
}
