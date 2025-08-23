package service

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

var jsonToReturn = `{
  "ts": 1755976374080,
  "tsj": 1755976365530,
  "date": "Aug 23rd 2025, 03:12:45 pm NY",
  "items": [
    {
      "curr": "BRL",
      "xauPrice": 18283.5602,
      "xagPrice": 210.8394,
      "chgXau": -12.1124,
      "chgXag": 2.0845,
      "pcXau": -0.0662,
      "pcXag": 0.9985,
      "xauClose": 18295.67258,
      "xagClose": 208.75486
    }
  ]
}`

type RoundTripFunc func(req *http.Request) *http.Response

func (t RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return t(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{Transport: fn}
}

func TestGetPrices(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
			Header:     make(http.Header),
		}
	})

	g := Gold{
		Items:  nil,
		Client: client,
	}

	_, err := g.GetPrices()

	if err != nil {
		t.Error(err)
	}
}
