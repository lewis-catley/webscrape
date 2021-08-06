package page

import (
	"net/url"

	"golang.org/x/net/html"
)

// Page store all relevant values found on a web page
type Page struct {
	URLChannel  chan *url.URL
	QuitChannel chan bool
	nodes       *html.Node
	urlsVisited []*url.URL
	err         error
	baseURL     *url.URL
}

// New returns a new instance of Page
func New(baseURL *url.URL) *Page {
	return &Page{
		URLChannel:  make(chan *url.URL),
		QuitChannel: make(chan bool),
		baseURL:     baseURL,
	}
}

// SetError set the err value on the Page
func (p *Page) SetError(err error) {
	p.err = err
}

// SetNodes set the nodes value on the Page
func (p *Page) SetNodes(n *html.Node) {
	p.nodes = n
}

// Print prints all the urls found on the page, or the error
func (p *Page) Print() []string {
	out := []string{}
	for _, u := range p.urlsVisited {
		if p.err == nil {
			out = append(out, u.String())
		}
	}
	return out
}
