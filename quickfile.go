package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/justinas/alice"
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
	h := http.FileServer(http.Dir("."))
	c := alice.New(loggingHandler, maxAgeHandler).Then(h)
	log.Fatal(http.ListenAndServe(":"+port, c))
}
