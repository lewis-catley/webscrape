package page

import (
	"net/url"
	"reflect"
	"testing"

	"golang.org/x/net/html"
)

func TestGetLinks(t *testing.T) {
	tests := []struct {
		name        string
		quitChannel bool
	}{
		{
			name:        "Channel should be set to true",
			quitChannel: true,
		},
	}

	for _, test := range tests {
		p := New(&url.URL{})
		go p.GetLinks()

		actual := <-p.QuitChannel
		if test.quitChannel != actual {
			t.Fatalf(`
%s, QuitChannel not set correctly.
Actual: %t.
Expected: %t.
`, test.name, test.quitChannel, <-p.QuitChannel)
		}
	}
}

func TestFindLink(t *testing.T) {
	tests := []struct {
		name     string
		node     *html.Node
		expected *url.URL
	}{
		{
			name: "Should not find link",
			node: &html.Node{
				Data: "html",
				Type: html.ElementNode,
			},
		},
		{
			name:     "Should not find link when node is nil",
			node:     nil,
			expected: nil,
		},
		{
			name: "Should ignore empty href values",
			node: &html.Node{
				Data: "a",
				Type: html.ElementNode,
				Attr: []html.Attribute{
					{
						Key: "href",
						Val: "",
					},
				},
			},
			expected: nil,
		},
		{
			name: "Should ignore refernces to the same page",
			node: &html.Node{
				Data: "a",
				Type: html.ElementNode,
				Attr: []html.Attribute{
					{
						Key: "href",
						Val: "#path",
					},
				},
			},
			expected: nil,
		},
		{
			name: "Should find a tag, and send url through channel",
			node: &html.Node{
				Data: "a",
				Type: html.ElementNode,
				Attr: []html.Attribute{
					{
						Key: "href",
						Val: "/path",
					},
				},
			},
			expected: &url.URL{
				Scheme: "https",
				Host:   "test.com",
				Path:   "/path",
			},
		},
	}

	for _, test := range tests {
		var actual *url.URL
		p := New(&url.URL{
			Scheme: "https",
			Host:   "test.com",
		})
		go p.findLink(test.node)
		actual = <-p.URLChannel

		if !reflect.DeepEqual(test.expected, actual) {
			t.Fatalf(`
%s exepected does not equal actual.
Expected: %#v
Actual: %#v
`, test.name, test.expected, actual)
		}
	}
}
