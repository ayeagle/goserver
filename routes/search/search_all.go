package search

import (
	"fmt"
	"net/http"
)

func SearchAllHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("search all data")
	fmt.Fprintf(w, "search all data")
}
