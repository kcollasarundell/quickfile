package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func loggingHandler(h http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, h)
}

func maxAgeHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
		h.ServeHTTP(w, r)
	})
}

func main() {
	port := "8080"
	if len(os.Args) >= 2 {
		port = os.Args[1]
	}
	fileHandler := http.FileServer(http.Dir("."))
	cacheControl := maxAgeHandler(fileHandler)
	logItAll := loggingHandler(cacheControl)
	log.Fatal(http.ListenAndServe(":"+port, logItAll))
}
