package search

import (
	"fmt"
	"net/http"
)

func SearchAllHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("search all data")
	fmt.Fprintf(w, "Get off your fat lazy ass Kevin and just write the shit down, so that I can understand it and discuss it with you. K thx bye.")
}
