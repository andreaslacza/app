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
	mux.Handle("/", middleware(handler()))

	srv := &http.Server{
		Addr: *httpAddr,
		Handler: mux,
	}
	log.Fatal(srv.ListenAndServe())
}

type User struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Session {
	ID int64 `json:"id"`
}

func parseTemplate(filename string) *template.Template {
	tmpl := template.Must(template.ParseFiles("templates/base.html"))

	path := filepath.Join("templates", filename)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("could not read template: %w", err))
	}
	template.Must(tmpl.New("body").Parse(string(b)))

	return tmpl.Lookup("base.html")
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
