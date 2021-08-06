package page

import (
	"net/url"

	"golang.org/x/net/html"
)

// GetLinks triggers a loop through p.nodes once finished passes a value to the Quit Channel
func (p *Page) GetLinks() {
	if p.nodes != nil {
		p.loopNodes(p.nodes)
	}
	p.nodes = nil
	p.QuitChannel <- true
}

func (p *Page) loopNodes(nodes *html.Node) {
	p.findLink(nodes)

	for child := nodes.FirstChild; child != nil; child = child.NextSibling {
		p.loopNodes(child)
	}

	return
}

func (p *Page) findLink(n *html.Node) {
	if n == nil {
		return
	}

	// Specifically look at a tags in a document
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				val := attr.Val
				if len(val) == 0 {
					continue
				}
				if val[0] == '#' {
					continue
				}
				u, err := url.Parse(val)
				if err != nil {
					continue
				}
				newURL := p.baseURL.ResolveReference(u)
				p.urlsVisited = append(p.urlsVisited, newURL)
				p.URLChannel <- newURL
			}
		}
	}
	return
}
