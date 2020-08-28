package main

import (
	"log"
	"url_shortener/internal/app/url_shortener"
)

// Main starts the application
func main() {
	log.Println("Starting...")
	urlShortener.Start()
}
