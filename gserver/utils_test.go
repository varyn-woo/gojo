package gserver_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
)

// gserverPost sends a POST request to the test gserver at endpoint with the given string body
// and handles all CORS headers so the request doesn't get rejected
func gserverPost(s *httptest.Server, endpoint string, body string) (*http.Response, error) {
	req, err := http.NewRequest(
		"POST",
		s.URL+endpoint,
		strings.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	// set CORS origin header so it doesn't get rejected
	req.Header.Add("Origin", "test")
	return s.Client().Do(req)
}

// gserverGet sends a GET request to the test gserver at endpoint
func gserverGet(s *httptest.Server, endpoint string) (*http.Response, error) {
	return s.Client().Get(s.URL + endpoint)
}
