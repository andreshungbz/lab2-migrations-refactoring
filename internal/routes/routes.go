// Filename: internal/routes/routes.go
// Registers all routes with handlers and applies all middleware.

package routes

import (
	"net/http"

	"github.com/andreshungbz/lab2-migrations-refactoring/internal/data"
	"github.com/andreshungbz/lab2-migrations-refactoring/internal/handlers"
	"github.com/andreshungbz/lab2-migrations-refactoring/internal/middleware"
)

// SetupRoutes registers all handlers to the router and applies all route-specific middleware.
func SetupRoutes(mux *http.ServeMux) {
	// Route-specific Middleware
	calculateHandler := middleware.TimingMiddleware(http.HandlerFunc(handlers.CalculateHandler))

	// Handler Registration (mux.Handle used for http.HandlerFunc handlers)
	mux.HandleFunc(data.Routes["home"].Path, handlers.HomeHandler)
	mux.HandleFunc(data.Routes["about"].Path, handlers.AboutHandler)
	mux.HandleFunc(data.Routes["contact"].Path, handlers.ContactHandler)
	mux.Handle(data.Routes["calculate"].Path, calculateHandler)
}

// SetupMux applies global middleware.
func SetupMux(mux *http.ServeMux) http.Handler {
	return middleware.LoggingMiddleware(mux) // apply logging middleware
}
