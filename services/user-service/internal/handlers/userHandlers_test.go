package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/models"
	"github.com/stretchr/testify/assert"
) 


func TestCreateUser(t *testing.T){
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	db = mockDB
	router := mux.NewRouter()
	router.HandleFunc("/users", CreateUser).Methods("POST")

	user := models.User{NickName: "jeanToad", Email: "jeanToad@gmail.com", Password: "JeanToad123@"}
	userJSON, _ := json.Marshal(user)
	request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	response := httptest.NewRecorder()

	mock.ExpectQuery(`INSERT INTO users\(nickname, email, password\) VALUES\(\$1, \$2, \$3\) RETURNING id`).
	WithArgs(user.NickName, user.Email, sqlmock.AnyArg()).
	WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	router.ServeHTTP(response, request)
	assert.Equal(t, http.StatusCreated, response.Code, "Expected response code to be 201")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
func TestGetUser(t *testing.T) {
    mockDB, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer mockDB.Close()
    db = mockDB 

    router := mux.NewRouter()
    router.HandleFunc("/users/{id}", GetUser).Methods("GET")

    user := models.User{
        NickName: "jeanToad",
        Email:    "jeanToad@gmail.com",
        IDToken:  "1",
    }

    mock.ExpectQuery(`SELECT email, nickname, id_token FROM users WHERE id = \$1`).
        WithArgs(user.IDToken).
        WillReturnRows(sqlmock.NewRows([]string{"email", "nickname", "id_token"}).
            AddRow(user.Email, user.NickName, user.IDToken))

    request, _ := http.NewRequest("GET", "/users/"+user.IDToken, nil)
    response := httptest.NewRecorder()

    router.ServeHTTP(response, request)

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }

    assert.Equal(t, http.StatusOK, response.Code, "Expected HTTP status 200")
    var returnedUser models.User
    json.NewDecoder(response.Body).Decode(&returnedUser)

    assert.Equal(t, user.Email, returnedUser.Email, "Email should match")
    assert.Equal(t, user.NickName, returnedUser.NickName, "Nickname should match")
    assert.Equal(t, user.IDToken, returnedUser.IDToken, "IDToken should match")
}
