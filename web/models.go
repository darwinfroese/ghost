package web

import (
	"net/http"
)

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
}

// Response is an object for storing expected responses
type Response struct {
	Count int
	Code  int
	Body  interface{}
}
