package main

import (
	"log"
	"net/http"

	"github.com/HazimEmam/JWTtutorial/database"
	"github.com/HazimEmam/JWTtutorial/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	cfg := database.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DBName:   "JWTtutorial",
	}

	database.InitDB(cfg)

	routes.Routes(router)
	routes.PagesRoutes(router)

	err := http.ListenAndServe(":8090", router)
	if err != nil {
		log.Panic("can't use Port 8090")
	}

}
