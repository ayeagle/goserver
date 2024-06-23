package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

var allowedOrigins = map[string]bool{
	"":                      true, // localhost won't send origin to localhost
	"http://localhost":      true,
	"http://localhost:8080": true,
}

const DEFAULT_REQUEST_PATH_TO_REMOVE = "/default/bennies"

func AddMiddlewareWrapper(router http.Handler) http.Handler {
	return stripAWSDefaultPathingMW(
		appendTrailingSlashMW(
			checkCORSMW(
				router)))
}

func isValidOrigin(origin string) bool {
	return allowedOrigins[origin]
}

func checkCORSMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		fmt.Printf("DEBUG: Testing origin: %s\n", origin)
		fmt.Printf("DEBUG: checkCORSMW path: %s\n", r.URL.Path)

		if isValidOrigin(origin) {
			fmt.Printf("OK ORIGIN\n")
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		} else {
			if origin != "" {
				fmt.Printf("FORBATTEN ORIGIN\n")
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			fmt.Printf("Likely local host connection, status OK\n")
		}

		// preflight handling
		if r.Method == http.MethodOptions {
			if isValidOrigin(origin) {
				fmt.Printf("OK ORIGIN -- preflight\n")
				w.WriteHeader(http.StatusOK)
				return
			}
			fmt.Printf("FORBATTEN ORIGIN -- preflight\n")
			w.WriteHeader(http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Should this be removed? TODO
func appendTrailingSlashMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "/") {
			r.URL.Path += "/"
		}
		next.ServeHTTP(w, r)
	})
}

func stripAWSDefaultPathingMW(next http.Handler) http.Handler {
	return http.StripPrefix(DEFAULT_REQUEST_PATH_TO_REMOVE, next)
}
