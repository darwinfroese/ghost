package web

import (
	"io"
	"net/http"
)

// Validator is a function that receives an HTTP response body
// and returns whether or not the response body was correct
type Validator func(body io.ReadCloser) bool

// Test is an object for a test scenario
type Test struct {
	Description string
	Tests       []RouteTest
}

// RouteTest is an object for storing a test case
type RouteTest struct {
	Description      string
	Route            string
	URL              string
	Method           string
	Body             []byte
	Handler          http.HandlerFunc
	ExpectedResponse Response
	Validator
}

// Response is an object for storing expected response information
type Response struct {
	Count int
	Code  int
	Body  interface{}
}
