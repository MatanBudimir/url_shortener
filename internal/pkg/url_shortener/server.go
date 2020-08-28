package shortener

import (
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
	"time"
)

var MemoryCache *cache.Cache

// server starts the web server
func serve() {
	MemoryCache = cache.New(5*time.Minute, 10*time.Minute)
	router := mux.NewRouter()

	router.HandleFunc("/", indexPage).Methods("GET", "POST")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../web/static"))))
	router.HandleFunc("/{uuid}", redirect).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}
