package controllers

import (
	"html/template"
	"log"
	"net/http"
	"web/models"
)

func UsersCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")

		models.UsersCreate(name, email, password)

		http.Redirect(w, r, "/users/", http.StatusFound)
	}

	tmp, err := template.ParseFiles("templates/html/create.html", "templates/html/layout.html")

	if err != nil {
		log.Fatal("Error initializing template")
	}

	tmp.Execute(w, nil)
}
