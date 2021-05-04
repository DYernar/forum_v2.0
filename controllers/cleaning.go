package controllers

import (
	"net/http"

	"forum_v2.0/db"
)

//Cleanup used to drop all tables
func Cleanup(w http.ResponseWriter, r *http.Request) {
	db.DropTables()
}
