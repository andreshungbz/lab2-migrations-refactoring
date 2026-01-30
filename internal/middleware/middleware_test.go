// Filename: internal/middleware/middleware_test.go
// - Tests for log content and if middleware was called.

package middleware

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/andreshungbz/lab2-migrations-refactoring/internal/handlers"
)

// UNIT TESTS

// TestLoggingMiddleware asserts the existence of the HTTP request method and path in the log.
func TestLoggingMiddleware(t *testing.T) {
	// Setup a temporary log buffer
	var buf bytes.Buffer
	prev := log.Writer()
	log.SetOutput(&buf)
	defer log.SetOutput(prev)

	// Setup handler with middleware
	next := http.HandlerFunc(handlers.HomeHandler)
	handler := TimingMiddleware(next)

	// Setup fake HTTP request and send to handler
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Assertion: Log string contain the HTTP request method and path
	out := buf.String()
	if !strings.Contains(out, "GET /") {
		t.Fatalf("Expected log to contain HTTP method and path, got %q", out)
	}
}

// TestTimingMiddleware asserts the existence of the '[TIME]' categorization in the log.
func TestTimingMiddleware(t *testing.T) {
	// Setup temporary log buffer
	var buf bytes.Buffer
	prev := log.Writer()
	log.SetOutput(&buf)
	defer log.SetOutput(prev)

	// Setup handler with middleware
	next := http.HandlerFunc(handlers.CalculateHandler)
	handler := TimingMiddleware(next)

	// Setup fake HTTP request and send to handler
	req := httptest.NewRequest("GET", "/calculate", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Assertion: Log string contains '[TIME]'
	out := buf.String()
	if !strings.Contains(out, "[TIME]") {
		t.Fatalf("Expected log to contain '[TIME]', got %q", out)
	}
}
