package userinfo

import(
	"net/http"
)

type UserInfo interface{
	 GetUserProfile( http.ResponseWriter,  *http.Request)
	 GetMicroServiceName( http.ResponseWriter,  *http.Request)
}