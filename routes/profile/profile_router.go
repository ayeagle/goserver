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
	"fmt"
	"net/http"
)

// RegisterRoutes registers the profile routes with the given ServeMux
func RegisterRoutes(mux *http.ServeMux) {
	fmt.Printf("Profile router hit\n")

	//  mux.HandleFunc("/", ProfileHandler)

	// mux.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Printf("Handler for / called. Path=%s\n", r.URL.Path)
	// 	fmt.Print("EMPTY handled")
	// 	ProfileHandler(w, r)
	// })
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler for / called. Path=%s\n", r.URL.Path)
		fmt.Print("one slash handled")

		ProfileHandler(w, r)
	})
	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler for /profile called. Path=%s\n", r.URL.Path)
		fmt.Print("pre slash handled")

		ProfileHandler(w, r)
	})
	mux.HandleFunc("/profile/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler for /profile called. Path=%s\n", r.URL.Path)
		fmt.Print("double slash handled")
		ProfileHandler(w, r)
	})
	mux.HandleFunc("profile/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler for /profile called. Path=%s\n", r.URL.Path)
		fmt.Print("post slash handled")
		ProfileHandler(w, r)
	})
}
