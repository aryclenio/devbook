package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, error := io.ReadAll(r.Body)
	if error != nil {
		log.Fatal(error)
	}
	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		log.Fatal(error)
	}
	db, error := database.Connect()
	if error != nil {
		log.Fatal(error)
	}

	repository := repositories.NewUserRepository(db)
	repository.Create(user)
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user"))
}
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user"))
}
