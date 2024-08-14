package handlers

import (
	"log"
	"net/http"
	"strconv"

	art "ascii-art-web/ascii-art"
)

// DownloadAsciiArt handles GET requests to download ASCII art as a .txt file.
func DownloadAsciiArt(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request on /download route\n", r.Method)

	if r.Method != http.MethodGet {
		renderErrorPage(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract filename from query parameters
	filename := r.URL.Query().Get("filename")
	input := r.URL.Query().Get("input")
	
	if filename == "" || input == "" {
		renderErrorPage(w, "Bad Request: Missing filename or input", http.StatusBadRequest)
		return
	}

	// Generate ASCII art
	result, err := art.AsciiArt(input, filename)
	if err != nil {
		log.Printf("Error generating ASCII art: %v\n", err)
		renderErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set headers for file download
	w.Header().Set("Content-Disposition", "attachment; filename=\"ascii-art.txt\"")
	w.Header().Set("Content-Type", "text/plain")
	contentLength := len(result)
	w.Header().Set("Content-Length", strconv.Itoa(contentLength))

	// Write the result to the response
	_, err = w.Write([]byte(result))
	if err != nil {
		log.Printf("Error writing response: %v\n", err)
		renderErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
