package home

import (
	"fmt"
	"net/http"
)

// HomeHandler handles requests to the root URL.
func HomeRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("home root data")
	fmt.Fprintf(w, "home data: root")
}
