// Filename: internal/handlers/handlers.go
// - Handler functions executed by the router.
package handlers

import (
	"fmt"
	"net/http"

	"github.com/andreshungbz/lab2-migrations-refactoring/internal/data"
)

// HomeHandler writes a welcome message and elaborates on the choice of semester project.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(data.Routes["home"].Text))
}

// AboutHandler writes information about myself (the student).
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(data.Routes["about"].Text))
}

// ContactHandler writes a student email address and their GitHub username.
func ContactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(data.Routes["contact"].Text))
}

// CalculateHandler writes a simple mathematical expression.
func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(fmt.Appendf(nil, "The result of 6 + 7 = %d!\n", 6+7))
}
