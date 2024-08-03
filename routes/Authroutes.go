package routes

import (
	"github.com/HazimEmam/JWTtutorial/controllers"
	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/register", controllers.Register).Methods("POST")

}
