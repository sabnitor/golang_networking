package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/translate", translate).Methods("POST")
	http.ListenAndServe("localhost:8080", r)
}
