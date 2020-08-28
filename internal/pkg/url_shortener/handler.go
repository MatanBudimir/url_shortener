package shortener

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jakehl/goid"
	"github.com/patrickmn/go-cache"
	"html/template"
	"log"
	"net/http"
	"url_shortener/internal/pkg/database"
)

// URL represents short URL information
type URL struct {
	Uuid, Url string
}

// IndexPage handles the "/" route page
// In case the request method is POST it creates a new URL
// Anything else returns the create url page
func indexPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()

		url := r.Form.Get("url")

		if url == "" || len(url) > 255 {
			http.Redirect(w, r, "/?error=Form validation failed", 301)
			return
		}

		data := URL{goid.NewV4UUID().String(), url}

		result := database.DB.Create(&data)

		log.Println(data)

		if result.Error != nil {
			http.Redirect(w, r, fmt.Sprintf("/?error=%s", result.Error.Error()), 301)
			return
		}

		MemoryCache.Set(data.Uuid, url, cache.NoExpiration)

		http.Redirect(w, r, "/?success=URL Created", 301)
	} else {
		t, err := template.ParseFiles("../../web/template/index.html")
		if err != nil {
			panic(err)
		}

		e := r.URL.Query().Get("error")
		success := r.URL.Query().Get("success")

		t.Execute(w, map[string]string{
			"error":   e,
			"success": success,
		})
	}
}

// redirect function gets the url from the cache or the database. In case it's not found, it returns redirect to "/"
func redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if value, err := MemoryCache.Get(vars["uuid"]); err {

		http.Redirect(w, r, fmt.Sprintf("%s", value), 308)
	} else {

		var data URL

		result := database.DB.Where("uuid  = ?", vars["uuid"]).First(&data)

		if result.Error != nil {
			http.Redirect(w, r, fmt.Sprintf("/?error=%s", result.Error.Error()), 301)
			return
		}

		http.Redirect(w, r, data.Url, 308)
	}
}
