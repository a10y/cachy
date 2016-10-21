// Implements server for cachy
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/andreweduffy/cachy/cache"
)

// server type
type server struct {
	cache *cache.Cache
}

func (s *server) Serve(port int) {
	// Setup an HTTP server and listen to POST requests for the given thing
	http.HandleFunc("/set/", func(resp http.ResponseWriter, req *http.Request) {
		// Handle setting the key
		path := req.URL.Path
		elems := strings.Split(path, "/") // "","set",key,value
		key := elems[2]
		val := elems[3]
		log.Printf("SET %v -> %v", key, val)
		s.cache.Add(key, val)
		fmt.Fprintf(resp, "success")
	})

	// fetching implementation
	http.HandleFunc("/get/", func(resp http.ResponseWriter, req *http.Request) {
		// Handle setting the key
		path := req.URL.Path
		elems := strings.Split(path, "/") // "","get",key
		key := elems[2]
		if v, err := s.cache.Lookup(key); err != nil {
			log.Printf("MISS %v", key)
			http.NotFound(resp, req)
			return
		} else {
			log.Printf("HIT %v", key)
			fmt.Fprintf(resp, v.(string))
		}
	})
	http.HandleFunc("/evict", func(resp http.ResponseWriter, req *http.Request) {
		log.Printf("Evicting...")
		s.cache.Evict()
		fmt.Fprintf(resp, "evicted")
	})
	addr := fmt.Sprintf("0.0.0.0:%v", port)
	fmt.Printf("Serving @ %v\n", addr)
	http.ListenAndServe(addr, nil)
}

func main() {
	s := &server{cache.New()}
	s.Serve(8080)
}
