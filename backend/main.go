package main

import (
	"log"
	"net/http"

	"lru-cache/cache"
	"lru-cache/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize cache
	lruCache := cache.NewLRUCache(1024)

	// Create router
	router := mux.NewRouter()

	// Define CORS middleware
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}

	// Use CORS middleware
	router.Use(corsMiddleware)

	// Define API routes
	router.HandleFunc("/cache", handlers.GetAllCacheHandler(lruCache)).Methods(http.MethodGet)
	router.HandleFunc("/cache/{key}", handlers.GetHandler(lruCache)).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/cache/{key}", handlers.SetHandler(lruCache)).Methods(http.MethodPost, http.MethodOptions)

	// Start server
	log.Fatal(http.ListenAndServe(":8080", router))
}
