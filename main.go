package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/krafthack/reed/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/krafthack/reed/reeds"
)

func handler(w http.ResponseWriter, r *http.Request) {
	repo, err := reeds.NewReedsRepo("localhost", "test", "reedstest")
	if err != nil {
		fmt.Fprintf(w, "Could not connect to mongodb")
		return
	}

	reed := reeds.Reed{"Mostly harmless"}

	err = repo.Store(&reed)

	if err != nil {
		fmt.Fprintf(w, "Could not save Mostly harmless")
	} else {
		fmt.Fprintf(w, "Mostly harmless was saved")
	}

}

func main() {
	port := os.Getenv("PORT")

	r := mux.NewRouter()
	r.HandleFunc("/", handler)

	http.Handle("/", r)

	http.ListenAndServe(":"+port, nil)
}
