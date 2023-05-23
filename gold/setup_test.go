package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"

	"fyne.io/fyne/v2/test"
)

var testApp Config

func TestMain(m *testing.M) {
	a := test.NewApp()
	testApp.App = a
	testApp.HTTPClient = client
	os.Exit(m.Run())
}

var jsonToReturn = `
{
	"ts": 1684853737933,
	"tsj": 1684853735541,
	"date": "May 23rd 2023, 10:55:35 am NY",
	"items": [
	{
	"curr": "USD",
	"xauPrice": 1969.7850,
	"xagPrice": 23.478,
	"chgXau": -1.755,
	"chgXag": -0.1205,
	"pcXau": -0.0891,
	"pcXag": -0.5106,
	"xauClose": 1969.785,
	"xagClose": 23.5985
	}
	]
	}
`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

var client = NewTestClient(func(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
		Header:     make(http.Header),
	}
})
