// API
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	WEBPORT = ":8080"
)

func main() {

	router := mux.NewRouter()

	http.Handle("/", router)

	router.HandleFunc("/home", homefunc)

	http.ListenAndServe(WEBPORT, nil)
}

func homefunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the home page")
}
