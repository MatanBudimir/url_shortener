package shortener

import (
	"url_shortener/internal/pkg/database"
)

type User struct {
	ID                  string
	NameFirst, NameLast string
}

// Configure starts the app
func Configure() {
	database.Configure()

	serve()
}
