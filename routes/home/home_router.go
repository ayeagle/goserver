// package home

// import (
// 	"goserver/routes"
// 	"net/http"
// )

// func init() {
// 	routes.RegisterHandlersForPackage("home", "./routes/home", HomePackage)
// }

// var HomePackage = struct {
// 	HomeHandler func(http.ResponseWriter, *http.Request)
// }{
// 	HomeHandler: HomeHandler,
// }

package home

import "net/http"

// RegisterRoutes registers the home routes with the given ServeMux
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", HomeRootHandler)
}
