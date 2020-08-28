package main
import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"net/http"
	"time"
	"github.com/jakehl/goid"
)

var MemoryCache *cache.Cache

func main() {
	MemoryCache = cache.New(5*time.Minute, 1*time.Minute)

	http.HandleFunc("/", test)
	http.HandleFunc("/test", find)
	http.ListenAndServe(":3000", nil)
}

func find(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("url")
	_, err := MemoryCache.Get(param)
	if !err {
		fmt.Fprintf(w, "Not found")
		return
	}

	fmt.Fprintf(w, "Found")
}

func test(w http.ResponseWriter, r *http.Request) {
	uid := goid.NewV4UUID()
	_, err := MemoryCache.Get(uid.String())
	if !err {
		MemoryCache.Set(uid.String(), 1, cache.NoExpiration)
	}

	fmt.Fprintf(w, fmt.Sprintf("%v", uid.String()))
}