package main

import (
	"log"
	"userinfo"
	"net/http"

)



func main() {

	http.HandleFunc("/user/profile",userinfo.ServiceInstance.GetUserProfile)
	http.HandleFunc("/microservice/name",userinfo.ServiceInstance.GetMicroServiceName)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
