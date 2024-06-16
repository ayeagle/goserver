// package profile

// import (
// 	"goserver/routes"
// 	"net/http"
// )

// func init() {
// 	routes.RegisterHandlersForPackage("profile", "./routes/profile", ProfilePackage)
// }

// var ProfilePackage = struct {
// 	ProfileHandler func(http.ResponseWriter, *http.Request)
// }{
// 	ProfileHandler: ProfileHandler,
// }

package profile

import (
	"net/http"
)

// RegisterRoutes registers the profile routes with the given ServeMux
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", ProfileHandler)
}
