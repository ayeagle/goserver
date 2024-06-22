// package search

// import (
// 	"net/http"

// 	"goserver/routes"
// )

// func init() {
// 	routes.RegisterHandlersForPackage("search", "./routes/search", SearchPackage)
// }

// // Include new routes within directory here
// var SearchPackage = struct {
// 	SearchAllHandler func(http.ResponseWriter, *http.Request)
// 	SearchOneHandler func(http.ResponseWriter, *http.Request)
// }{
// 	SearchAllHandler: SearchAllHandler,
// 	SearchOneHandler: SearchOneHandler,
// }

package search

import (
	"net/http"
)

// RegisterRoutes registers the search routes with the given ServeMux
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", SearchRootHandler)
	mux.HandleFunc("/all/", SearchAllHandler)
	mux.HandleFunc("/one/", SearchOneHandler)
}
