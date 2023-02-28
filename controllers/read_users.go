package controllers

import (
	"html/template"
	"log"
	"net/http"
	"web/models"
)

func ReadUsers(w http.ResponseWriter, r *http.Request) {

	tmp, err := template.ParseFiles("templates/html/readUsers.html", "templates/html/layout.html")

	if err != nil {
		log.Fatal("Error initializing template readUsers!")
	}

	usuarios := models.ReadUser()

	tmp.Execute(w, usuarios)
}
