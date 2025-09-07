package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler).Methods("GET")

	fmt.Println("Running server")
	log.Fatal(http.ListenAndServe(":9000", r))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Index is showingg")
	w.Write([]byte("<h1>Index Page</h1>"))
}
