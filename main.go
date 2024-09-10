package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// header fucntion
func noCacheFileServer(dir string) http.Handler {
	fs := http.FileServer(http.Dir(dir))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//headers
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		w.Header().Set("Surrogate-Control", "no-store")

		fs.ServeHTTP(w, r)
	})
}

func main() {
	flag.String("help", "", "display help")
	port := flag.String("port", "8080", "Port to serve on")
	flag.Parse()

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}

	fileServer := noCacheFileServer(dir)

	http.Handle("/", fileServer)

	fmt.Printf("Serving files from %s on http://localhost:%s with no-cache headers\n", dir, *port)
	err = http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
