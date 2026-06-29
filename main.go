package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var httpAddr = flag.String("http", "localhost:6060", "HTTP service address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.Handle("/", handler())

	srv := &http.Server{
		Addr: *httpAddr,
		Handler: mux,
	}
	log.Fatal(srv.ListenAndServe())
}

type User struct {
	ID int64
}

func handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}
