package web

import (
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

type errorClient struct{}

func (ec *errorClient) Get(url string) (*http.Response, error) {
	return nil, errors.New("The Get failed")
}

type successClient struct{}

func (sc *successClient) Get(url string) (*http.Response, error) {
	return &http.Response{
		Body: http.NoBody,
	}, nil
}

func TestGetWebsiteContent(t *testing.T) {
	tests := []struct {
		name        string
		input       url.URL
		client      HTTPClient
		expected    *[]byte
		errExpected bool
	}{
		{
			name:        "Get fails",
			input:       url.URL{},
			client:      &errorClient{},
			errExpected: true,
		},
		{
			name: "Get success",
			input: url.URL{
				Host: "bbc.co.uk",
			},
			client:   &successClient{},
			expected: &[]byte{},
		},
	}

	for _, test := range tests {
		Client = test.client
		actual, err := GetURLContent(test.input)
		if test.errExpected != (err != nil) {
			t.Fatalf(`
%s, failed with unexpected error
Actual %v.
Expected %t.
`, test.name, err, test.errExpected)
		}

		if !reflect.DeepEqual(test.expected, actual) {
			t.Fatalf(`
%s, failed with different responses.
Actual %v.
Expected %v.
`, test.name, actual, test.expected)
		}
	}
}
