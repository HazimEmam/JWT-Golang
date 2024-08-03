package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/HazimEmam/JWTtutorial/database"
	"github.com/HazimEmam/JWTtutorial/models"
	"github.com/HazimEmam/JWTtutorial/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("can't Marshal The user"))
		return
	}
	var existingUser models.User
	database.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Email Not existes"))
		return
	}
	ok := utils.CompareHashPassword(user.Password, existingUser.Password)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Invaild Password"))
		return
	}

	token, err := utils.GenerateToken(existingUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Field when generate token"))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(time.Minute * 15),
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("can't Marshal The user"))
		return
	}
	var existingUser models.User
	database.DB.Where("email = ?", user.Email).First(&existingUser)
	if existingUser.ID != 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Email existes before"))
		return
	}
	user.Password, err = utils.GenerateHashPassword(user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error when Hash the password"))
		return
	}
	database.DB.Create(&user)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User registers successfully"))

}
