package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	DOB   string `json:"dob"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/user/profile", nil)
	res, _ := client.Do(req)
	defer res.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	data := User{}
	json.Unmarshal(bodyBytes, &data)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GetMicroService(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/microservice/name", nil)
	res, _ := client.Do(req)

	defer res.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	bodyString := string(bodyBytes)
	fmt.Fprintf(w, bodyString)

}
