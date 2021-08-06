package web

import (
	"fmt"
	"io"
	"net/url"

	"golang.org/x/net/html"
)

// ParseReaderHTMLNode will attempt to parse a reader into *html.Node
func ParseReaderHTMLNode(r io.Reader) (*html.Node, error) {
	n, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return n, nil
}

// GenerateURLNoParams will return a raw url without any path parameters
func GenerateURLNoParams(u *url.URL) string {
	return fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, u.Path)
}
