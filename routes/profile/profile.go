package profile

import (
	"fmt"
	"net/http"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("this is the proflie data lmao")

	// Set a custom header
	w.Header().Set("Content-Type", "text/plain")

	// Write the HTTP status code
	w.WriteHeader(http.StatusOK)

	// Write the response body
	fmt.Fprintln(w, "Hello, World!")

	fmt.Printf("this is the proflie data lmao")
	fmt.Fprintf(w, "profile data")
	fmt.Printf("After I guess?")

}
