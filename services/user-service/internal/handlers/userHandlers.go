package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

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