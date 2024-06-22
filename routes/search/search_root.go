package search

import (
	"fmt"
	"net/http"
)

func SearchRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("search root data")
	fmt.Fprintf(w, "search data: root")
}
