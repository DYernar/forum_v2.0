package controllers

import (
	"html/template"
	"net/http"

	"forum_v2.0/db"
	"forum_v2.0/model"
)

// MainPage controller
func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if r.Method == "GET" {
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
			response.Posts = db.GetAllPosts()
			t := template.Must(template.New("index").ParseFiles("static/index.html", "static/header.html", "static/footer.html"))
			t.Execute(w, response)
		} else {
			BadRequest(w, r, r.Method+" is not allowed")
		}
	} else {
		PageNotFound(w, r)
	}

}
