package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/krafthack/reed/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/krafthack/reed/reeds"
)

func handler(w http.ResponseWriter, r *http.Request) {
	repo, err := reeds.NewReedsRepo("localhost", "test", "reedstest")
	if err != nil {
		fmt.Fprintf(w, "Could not connect to mongodb")
		return
	}

	reed := reeds.Reed{
		"Rarely say yes to feature requests",
		"Here’s a simple set of Yes/No questions that you can quickly answer before you add another item to your product roadmap.  Saying yes to a feature request – whether it’s a to an existing customer, a product enquiry, a teammate, or a manager – is immediately rewarding.",
		"https://blog.intercom.io/wp-content/uploads/2014/11/Rarely-Say-Yes-984.jpg",
		"https://blog.intercom.io/rarely-say-yes-to-feature-requests/?utm_medium=email&utm_source=email&utm_campaign=say-no-email",
		time.Now().Add(-time.Hour * 24 * 3),
		time.Now(),
	}

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
