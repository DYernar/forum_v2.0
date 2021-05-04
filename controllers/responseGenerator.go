package controllers

import (
	"net/http"

	"forum_v2.0/db"
	"forum_v2.0/model"
)

//GetResponse adds user details to the response
func GetResponse(r *http.Request) model.Response {
	var response model.Response
	uname, loggedIn := IsAuthorized(r)
	if loggedIn {
		user, err := db.GetUserByEmail(uname)
		if err != nil {
			user = db.GetUserByName(uname)
		}
		response.User = user
	}
	response.LoggedIn = loggedIn
	return response
}
