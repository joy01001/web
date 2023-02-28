package controllers

import (
	"net/http"
	"web/models"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	models.DeleteUser(id)

	http.Redirect(w, r, "/users/", http.StatusFound)
}
