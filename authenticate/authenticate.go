package authenticate

import (
	"fmt"
	"net/http"
)

func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		client := &http.Client{}
		req, _ := http.NewRequest("GET", "http://localhost:9001/auth", nil)

		req.Header.Set("UserName", r.Header.Get("UserName"))
		res, _ := client.Do(req)

		if res.StatusCode == http.StatusOK {
			endpoint(w, r)

		} else {
		fmt.Fprintf(w,"Unauthorized")

		}

	})
}
