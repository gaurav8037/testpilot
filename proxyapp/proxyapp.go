package main

import (
	auth "authenticate"
	"log"
	"net/http"
	usr "user"
)

func handleRequests() {
	http.Handle("/user/profile", auth.IsAuthorized(usr.GetUserProfile))

	http.HandleFunc("/microservice/name", usr.GetMicroService)

	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	handleRequests()
}
