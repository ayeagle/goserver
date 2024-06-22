package search

import (
	"fmt"
	"net/http"
)

func SearchOneHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("search one data")
	fmt.Fprintf(w, "search one data")
}
