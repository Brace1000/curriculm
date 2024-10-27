package main

import (
	"log"
	"net/http"
	"cur/handlers"
	
)

func main() {
	// Map URLs to their respective page handlers
	http.HandleFunc("/", handlers.HomeHandler)
	
	// // Serve static files (CSS, images)
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static", http.StripPrefix("/static", fileServer))

	// Start the server
	log.Println("Server running at http://localhost:8083")
	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}