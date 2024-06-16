package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"goserver/middleware"
	"goserver/routes/home"
	"goserver/routes/profile"
	"goserver/routes/search"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	env := os.Getenv("ENV")
	var serverURL string

	if env == "PROD" {
		serverURL = os.Getenv("PROD")
	} else if env == "LOCAL" {
		serverURL = os.Getenv("LOCAL")
	} else {
		log.Fatalf("Could not start server: \n%s\n", "No valid server URL chosen.")
	}

	mainRouter := createRouters()
	wrappedRouter := middleware.CORSMiddleware(mainRouter)

	fmt.Printf("Starting server on %s\n", serverURL)
	if err := http.ListenAndServe(serverURL, wrappedRouter); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

func createRouters() *http.ServeMux {

	// primary router
	mainRouter := http.NewServeMux()

	// subrouters
	homeRouter := http.NewServeMux()
	profileRouter := http.NewServeMux()
	searchRouter := http.NewServeMux()

	// register primary subroutes
	home.RegisterRoutes(homeRouter)
	profile.RegisterRoutes(profileRouter)
	search.RegisterRoutes(searchRouter)

	// register subrouters with the main router
	mainRouter.Handle("/", homeRouter)

	mainRouter.Handle("/home/", http.StripPrefix("/home", homeRouter))
	mainRouter.Handle("/profile/", http.StripPrefix("/profile", profileRouter))
	mainRouter.Handle("/search/", http.StripPrefix("/search", searchRouter))

	return mainRouter
}
