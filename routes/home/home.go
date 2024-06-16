package home

import (
	"fmt"
	"net/http"
)

// HomeHandler handles requests to the root URL.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home page data")
}
