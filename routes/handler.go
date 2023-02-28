package routes

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"web/controllers"
)

func Handler() {

	r := mux.NewRouter()

	usersRouter := r.PathPrefix("/users").Subrouter()

	usersRouter.HandleFunc("/", controllers.ReadUsers).Methods("GET")
	usersRouter.HandleFunc("/create", controllers.UsersCreate).Methods("GET", "POST")
	usersRouter.HandleFunc("/delete", controllers.DeleteUser).Methods("GET")
	usersRouter.HandleFunc("/update", controllers.UpdateUser).Methods("GET", "POST")

	//Port redirects.
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	log.Println("Server Running...\nhttp://localhost:8080")

	//Permissions.
	handler := cors.AllowAll().Handler(usersRouter)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
