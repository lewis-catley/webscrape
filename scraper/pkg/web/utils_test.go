package web

import (
	"errors"
	"io"
	"net/url"
	"strings"
	"testing"
)

func TestGenerareURLNoParams(t *testing.T) {
	tests := []struct {
		name     string
		u        *url.URL
		expected string
	}{
		{
			name: "Generates correct url, with path params",
			u: &url.URL{
				Scheme:   "https",
				Host:     "test.com",
				Path:     "/test",
				RawQuery: "test=123",
			},
			expected: "https://test.com/test",
		},
		{
			name: "Generates correct url, without path params",
			u: &url.URL{
				Scheme: "https",
				Host:   "test.com",
				Path:   "/test",
			},
			expected: "https://test.com/test",
		},
	}

	for _, test := range tests {
		actual := GenerateURLNoParams(test.u)
		if actual != test.expected {
			t.Fatalf(`
%s Has faied unexpected value.
Expected: %s.
Actual: %s.`, test.name, test.expected, actual)
		}
	}
}

type DummyReader struct{ Error error }

func (dr *DummyReader) Read([]byte) (int, error) {
	return 0, dr.Error
}

func TestParseHTML(t *testing.T) {

	tests := []struct {
		name        string
		input       io.Reader
		shouldError bool
	}{
		{
			name:  "Valid Reader",
			input: strings.NewReader("<html></html>"),
		},
		{
			name: "Invalid Reader",
			input: &DummyReader{
				Error: errors.New("This is a dummy error"),
			},
			shouldError: true,
		},
	}

	for _, test := range tests {
		_, err := ParseReaderHTMLNode(test.input)
		if err != nil && !test.shouldError {
			t.Fatalf(`
%s errored unexpectedly.
Failed with error: %v`, test.name, err)
		}

		if err == nil && test.shouldError {
			t.Fatalf("%s did not error", test.name)
		}

	}
}
