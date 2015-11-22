package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	port := os.Getenv("PORT")

	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	http.Handle("/", r)

	http.ListenAndServe(":"+port, nil)
}
