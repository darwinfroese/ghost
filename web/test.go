package web

import (
	"testing"
)

// TestRequests tests web requests
func TestRequests(t *testing.T, scenario Test) {
	t.Log(scenario.Description)

	for _, test := range scenario.Tests {
		t.Log(test.Description)

		r, w := buildHTTP(test.Method, test.URL, nil)
		test.Handler(w, r)

		if w.Code != test.ExpectedResponse.Code {
			t.Errorf("[Test Failed] :: Expected status code %d but got status code %d", test.ExpectedResponse.Code, w.Code)
		}
	}
}
