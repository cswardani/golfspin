package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}
