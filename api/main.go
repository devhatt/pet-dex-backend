package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	port := ":3000"
	log.Printf("🚀 Server is running on http://localhost%v", port)

	http.ListenAndServe(port, r)
}
