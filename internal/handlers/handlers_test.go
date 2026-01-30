// Filename: internal/handlers/handlers_test.go
// - HTTP test requests for response status and body.
package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andreshungbz/lab2-migrations-refactoring/internal/data"
)

// Unit Tests

func TestHomeHandler(t *testing.T) {
	assertResponse(t, data.Routes["home"].Path, HomeHandler, data.Routes["home"].Text)
}

func TestAboutHandler(t *testing.T) {
	assertResponse(t, data.Routes["about"].Path, AboutHandler, data.Routes["about"].Text)
}

func TestContactHandler(t *testing.T) {
	assertResponse(t, data.Routes["contact"].Path, ContactHandler, data.Routes["contact"].Text)
}

func TestCalculateHandler(t *testing.T) {
	assertResponse(t, data.Routes["calculate"].Path, CalculateHandler, data.Routes["calculate"].Text)
}

// Helper Functions

// assertHTTP200 ensures that the HTTP response status code is 200 OK.
func assertHTTP200(t *testing.T, got int) {
	t.Helper()
	if got != http.StatusOK {
		t.Errorf("got %v, expected %v", got, http.StatusOK)
	}
}

// assertResponseBody ensures that the HTTP response body content matches an expected string.
func assertResponseBody(t *testing.T, got string, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("got %v, expected %v", got, expected)
	}
}

// assertResponse sends a test HTTP request and asserts the response status code and body.
// The passed handler function is implicitly converted to http.HandlerFunc.
func assertResponse(t *testing.T, route string, handler http.HandlerFunc, expectedBody string) {
	t.Helper()
	req := httptest.NewRequest("GET", route, nil) // HTTP request
	rr := httptest.NewRecorder()                  // records HTTP response
	handler.ServeHTTP(rr, req)                    // send test request to handler

	// Assertions
	assertHTTP200(t, rr.Code)
	assertResponseBody(t, rr.Body.String(), expectedBody)
}
