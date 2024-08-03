package routes

import (
	"github.com/HazimEmam/JWTtutorial/controllers"
	"github.com/gorilla/mux"
)

func PagesRoutes(r *mux.Router) {
	r.HandleFunc("/home", controllers.HomePage).Methods("GET")
	r.HandleFunc("/premium", controllers.PremiumPage).Methods("GET")
}
