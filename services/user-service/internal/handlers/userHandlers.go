package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/database"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/models"
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


}
func GetUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	user := models.User{UserID: "1", Name: "kevin", Email: "kevinv@gmail.com"} // fake data

	if id == "1"{
		json.NewEncoder(w).Encode(user)
	}
}