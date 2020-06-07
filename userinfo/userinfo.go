package userinfo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	DOB   string `json:"dob"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Info struct {
	a string
}

var ServiceInstance = Info{
	"test",
}

func (i Info) GetUserProfile(w http.ResponseWriter, r *http.Request) {

	// fetch data from database , hardcoded for now
	w.Header().Set("Content-Type", "application/json")
	user := User{
		Name:  "Gaurav Soni",
		DOB:   "18/09/1994",
		Email: "g.ymca8037@gmail.com",
		Phone: "8743883506",
	}

	json.NewEncoder(w).Encode(user)
}

func (i Info) GetMicroServiceName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "user-microservice")
}
