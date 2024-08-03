package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/HazimEmam/JWTtutorial/utils"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("There is no token"))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	claims, err := utils.VerifyToken(c.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid Token"))
		return
	}

	marshal , _ := json.Marshal(claims)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshal)
}
func SalaryPage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("There is no token"))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	claims, err := utils.VerifyToken(c.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invaild token"))
		return
	}

	if claims.Role == "employee" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("This page isn't for employee"))
		return
	}

	marshal , _ := json.Marshal(claims)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshal)
}
