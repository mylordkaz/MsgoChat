package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/database"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/models"
	"github.com/mylordkaz/MsgoChat/services/user-service/pkg/hash"
)


var db *sql.DB

func init(){
	db = database.ConnectDB()
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	hashedPassword, err := hash.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

    // Insert user into the database
    err = db.QueryRow("INSERT INTO users(username, email, password) VALUES($1, $2, $3) RETURNING id", user.NickName, user.Email, user.Password).Scan(&user.IDToken)
    if err != nil {
        http.Error(w, "Error creating user", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)

}
func GetUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	var user models.User

	err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.IDToken)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
func UpdateUser(w http.ResponseWriter, r *http.Request){

}
func DeleteUser(w http.ResponseWriter, r *http.Request){

}