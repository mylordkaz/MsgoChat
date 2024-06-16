package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

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

	hashedPassword, err := hash.HashPassword(user.PasswordHash)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.PasswordHash = hashedPassword
	user.Provider = "local"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

    // Insert user into the database
    err = db.QueryRow("INSERT INTO users(name, email, password, provider, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6) RETURNING id", user.Name, user.Email, user.PasswordHash).Scan(&user.ID)
    if err != nil {
        http.Error(w, "Error creating user", http.StatusInternalServerError)
        return
    }
	w.Header().Set("content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)

}
func GetUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	var user models.User

	err := db.QueryRow("SELECT email, nickname, id_token FROM users WHERE id = $1", id).Scan(&user.Email, &user.Name, &user.AccessToken)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
func UpdateUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}
	_, err := db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", user.Name, user.Email, id)
	if err != nil{
		http.Error(w, "Failed updated user"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)

}
func DeleteUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil{
		http.Error(w, "Failed to delete user:"+err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "User deleted successfully"})

}