package search

import (
	"fmt"
	"net/http"
)

func SearchOneHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "search one data")
}
