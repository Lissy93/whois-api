package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/someshkar/whois-api/api"
)

func main() {
	// Handle single domain route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/")
		if path != "" && path != "multi" {
			api.MainHandler(w, r) // Pass the domain directly to your handler
		} else {
			// Handle root or incorrect path
			http.NotFound(w, r)
		}
	})

	// Handle multiple domains route
	http.HandleFunc("/multi", api.MultiHandler)

	// Choose the port to start server on
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	serverAddress := fmt.Sprintf(":%s", port)
	log.Printf("Starting server on %s", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, nil))

}
