package routes

import (

	"github.com/HazimEmam/JWTtutorial/controllers"
	"github.com/HazimEmam/JWTtutorial/middleware"
	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/login", middleware.IsAuthorized(controllers.Login)).Methods("POST")
	r.HandleFunc("/register", controllers.Register).Methods("POST")
}
