package search

import (
	"fmt"
	"net/http"
)

func SearchAllHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "search all data")
}
