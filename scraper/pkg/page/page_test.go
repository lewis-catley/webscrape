package page

import (
	"errors"
	"net/url"
	"reflect"
	"testing"

	"golang.org/x/net/html"
)

func TestSetError(t *testing.T) {
	tests := []struct {
		name   string
		err    error
		expErr bool
	}{
		{
			name:   "Correct error is set",
			err:    errors.New(""),
			expErr: true,
		},
		{
			name:   "Error is not set",
			expErr: false,
		},
	}

	for _, test := range tests {
		p := New(&url.URL{})
		p.SetError(test.err)

		if test.expErr != (p.err != nil) {
			t.Fatalf(`
%s, failed with unexpected error
Actual %v.
Expected %t.
`, test.name, p.err, test.expErr)
		}
	}
}

func TestSetNodes(t *testing.T) {
	tests := []struct {
		name  string
		nodes *html.Node
	}{
		{
			name: "Nodes should match passed",
			nodes: &html.Node{
				Type: html.ElementNode,
			},
		},
		{
			name:  "Nodes should be nil",
			nodes: nil,
		},
	}

	for _, test := range tests {
		p := New(&url.URL{})
		p.SetNodes(test.nodes)

		if !reflect.DeepEqual(test.nodes, p.nodes) {
			t.Fatalf(`
%s, failed with expected not matching actual.
Actual %v.
Expected %v.
`, test.name, p.nodes, test.nodes)
		}
	}
}
