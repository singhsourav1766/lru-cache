package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"lru-cache/cache"

	"github.com/gorilla/mux"
)

func GetAllCacheHandler(lruCache *cache.LRUCache) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Get all cache contents from the LRU cache
        cacheContents := lruCache.GetAll()

        // Encode cache contents to JSON
        jsonContents, err := json.Marshal(cacheContents)
        if err != nil {
            http.Error(w, "Failed to encode cache contents to JSON", http.StatusInternalServerError)
            return
        }

        // Set response headers
        w.Header().Set("Content-Type", "application/json")
        
        // Write JSON response
        w.WriteHeader(http.StatusOK)
        w.Write(jsonContents)
    }
}

// GetHandler handles GET requests to retrieve a value from the cache
func GetHandler(lruCache *cache.LRUCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		key := params["key"]

		value, found := lruCache.Get(key)
		if found {
			json.NewEncoder(w).Encode(value)
		} else {
			http.Error(w, "Key not found", http.StatusNotFound)
		}
	}
}

// SetHandler handles POST requests to set a value in the cache
func SetHandler(lruCache *cache.LRUCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		key := params["key"]

		var value interface{}
		if err := json.NewDecoder(r.Body).Decode(&value); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Set expiration time for key (5 seconds)
		lruCache.Set(key, value, 5*time.Second)

		// Return a JSON response with a success message
		jsonResponse := map[string]string{"message": "Cache item set successfully"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(jsonResponse)
	}
}
