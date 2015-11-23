package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"encoding/json"

	"github.com/krafthack/reed/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/krafthack/reed/reeds"
)

type IftttReeds struct {
	Title    string `json:"title"`
	URL      string `json:"url"`
	Excerpt  string `json:"excerpt"`
	ImageURL string `json:"imageUrl"`
	Tags     string `json:"tags"`
	AddedAt  string `json:"addedAt"`
}

func parse(body io.Reader) (*IftttReeds, error) {
	decoder := json.NewDecoder(body)
	var i IftttReeds
	err := decoder.Decode(&i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func handler(w http.ResponseWriter, r *http.Request) {

	body, err := parse(r.Body)

	if err != nil {
		fmt.Fprintln(w, "Could not parse input")
		return
	}

	repo, err := reeds.NewReedsRepo("localhost", "test", "reedstest")
	if err != nil {
		fmt.Fprintf(w, "Could not connect to mongodb")
		return
	}

	fmt.Printf("%v", body)

	reed := reeds.Reed{
		body.Title,
		body.Excerpt,
		body.ImageURL,
		body.URL,
		time.Now(),
		time.Now(),
	}

	err = repo.Store(&reed)

	if err != nil {
		fmt.Fprintf(w, "Could not save %v", body.Title)
	} else {
		fmt.Fprintf(w, "%v was saved", body.Title)
	}

}

func main() {
	port := os.Getenv("PORT")

	r := mux.NewRouter()
	r.HandleFunc("/", handler)

	http.Handle("/", r)

	http.ListenAndServe(":"+port, nil)
}
