package shortener

import (
	"url_shortener/internal/pkg/database"
)

// Configure starts the app
func Configure() {
	database.Configure()

	serve()
}
