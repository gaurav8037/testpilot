package main

import (
	"fmt"
	"log"
	"net/http"
)

func authenticate(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Username") == "Gaurav" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "authorized")
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func main() {
	http.HandleFunc("/auth", authenticate)
	log.Fatal(http.ListenAndServe(":9001", nil))
}
