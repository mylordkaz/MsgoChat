package test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/handlers"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/models"

	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *mux.Router){
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	
	handlers.SetDB(mockDB)
	router := mux.NewRouter()

	return mockDB, mock, router

}

func TestUserHandlers(t *testing.T) {
    t.Run("CreateUser", func(t *testing.T) {
        mockDB, mock, router := setupTest(t)
        defer mockDB.Close()

        router.HandleFunc("/users", handlers.CreateUser).Methods("POST")

        user := models.User{Name: "jeanToad", Email: "jeanToad@gmail.com", PasswordHash: "JeanToad123@"}
    userJSON, _ := json.Marshal(user)
    request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
    response := httptest.NewRecorder()

    // Update the mock expectation to match the actual query
    mock.ExpectQuery(`INSERT INTO users\(name, email, password, provider, created_at, updated_at\) VALUES\(\$1, \$2, \$3, \$4, \$5, \$6\) RETURNING id`).
        WithArgs(user.Name, user.Email, sqlmock.AnyArg(), "local", sqlmock.AnyArg(), sqlmock.AnyArg()).
        WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

    router.ServeHTTP(response, request)

    t.Logf("Response Body: %s", response.Body.String())

    assert.Equal(t, http.StatusCreated, response.Code, "Expected response code to be 201, got %d", response.Code)

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }

    var createdUser models.User
    err := json.NewDecoder(response.Body).Decode(&createdUser)
    if err != nil {
        t.Errorf("Error decoding response body: %v", err)
    }

    assert.Equal(t, user.Name, createdUser.Name, "Created user name should match input")
    assert.Equal(t, user.Email, createdUser.Email, "Created user email should match input")
    assert.NotEmpty(t, createdUser.ID, "Created user should have an ID")
    })

    // Add similar test cases for GetUser, UpdateUser, ...
    
	t.Run("GetUser", func(t *testing.T) {
        mockDB, mock, router := setupTest(t)
        defer mockDB.Close()

        router.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")

        user := models.User{
            
            Name:  "jeanToad",
            Email: "jeanToad@gmail.com",
        }

        mock.ExpectQuery(`SELECT email, name, id FROM users WHERE id = \$1`).
            WithArgs("1").
            WillReturnRows(sqlmock.NewRows([]string{"email", "name", "id"}).
                AddRow(user.Email, user.Name, user.ID))

        request, _ := http.NewRequest("GET", "/users/1", nil)
        response := httptest.NewRecorder()

        router.ServeHTTP(response, request)

        if err := mock.ExpectationsWereMet(); err != nil {
            t.Errorf("there were unfulfilled expectations: %s", err)
        }

        assert.Equal(t, http.StatusOK, response.Code, "Expected HTTP status 200")

        var returnedUser models.User
        err := json.NewDecoder(response.Body).Decode(&returnedUser)
        if err != nil {
            t.Errorf("Error decoding response body: %v", err)
        }

        assert.Equal(t, user.Email, returnedUser.Email, "Email should match")
        assert.Equal(t, user.Name, returnedUser.Name, "Name should match")
        assert.Equal(t, user.ID, returnedUser.ID, "ID should match")
    })

	t.Run("updateUser", func(t *testing.T) {
		mockDB, mock, router := setupTest(t)
    	defer mockDB.Close()
    
    	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")

		updatedUser := models.User{
			Name: "updatedName",
			Email: "updated@email.com",

		}
		updatedUserJSON, _ := json.Marshal(updatedUser)

		request, _ := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(updatedUserJSON))
		request = mux.SetURLVars(request, map[string]string{"id": "1"})
		response := httptest.NewRecorder()

		mock.ExpectExec(`UPDATE users SET name = \$1, email = \$2 WHERE id = \$3`).
		WithArgs(updatedUser.Name, updatedUser.Email, "1").
		WillReturnResult(sqlmock.NewResult(0, 1))

		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code, "Expected 200 code")

		var returnedUser models.User
		json.NewDecoder(response.Body).Decode(&returnedUser)
		assert.Equal(t, updatedUser.Name, returnedUser.Name, "Name should match")
		assert.Equal(t, updatedUser.Email, returnedUser.Email, "Email should match")

		if err := mock.ExpectationsWereMet(); err != nil {
        t	.Errorf("there were unfulfilled expectations: %s", err)
    	}
	})

}