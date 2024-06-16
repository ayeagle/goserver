package middleware

import (
	"fmt"
	"net/http"
)

var allowedOrigins = map[string]bool{
	"":                      true, // localhost won't send origin to localhost
	"http://localhost":      true,
	"http://localhost:8080": true,
	"https://djrabid.com":   true,
}

func isValidOrigin(origin string) bool {
	return allowedOrigins[origin]
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		origin := r.Header.Get("Origin")

		if isValidOrigin(origin) {
			fmt.Printf("OK ORIGIN")
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		} else {
			if origin != "" {
				fmt.Printf("FORBATTEN ORIGIN")
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			fmt.Printf("Likely local host connection, status OK\n")
		}

		// preflight handling
		if r.Method == http.MethodOptions {
			if isValidOrigin(origin) {
				fmt.Printf("OK ORIGIN")
				w.WriteHeader(http.StatusOK)
				return
			}
			fmt.Printf("FORBATTEN ORIGIN")
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
