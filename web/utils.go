package web

import (
	"bytes"
	"net/http"
	"net/http/httptest"
)

// buildHTTP creates a request and a response writer
func buildHTTP(method, url string, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder) {
	r := buildRequest(method, url, body)
	w := buildWriter()

	return r, w
}

// buildWriter builds a response buildWriter
func buildWriter() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
}

// buildRequest builds a request for testing routes
func buildRequest(method, url string, body *bytes.Buffer) *http.Request {
	if body == nil {
		return httptest.NewRequest(method, url, nil)
	}

	return httptest.NewRequest(method, url, body)
}
