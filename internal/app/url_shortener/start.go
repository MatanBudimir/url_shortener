package urlshortener

import (
	"fmt"
	"url_shortener/internal/pkg/url_shortener"
)

// Start calls the necessary function that starts the app
func Start() {
	fmt.Println("Starting")
	shortener.Configure()
}
