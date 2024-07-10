package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/mylordkaz/MsgoChat/services/auth-service/internal/models"
)


type Database struct {
	*sql.DB
}

func Newdatabase(connectionStr string) (*Database, error) {
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// verify connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	return &Database{DB: db}, nil
}

func (d *Database) Close() error {
	return d.DB.Close()
}

func (d *Database) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (username, password, email)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	err := d.QueryRow(query, user.Username, user.Password, user.Email).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}

func (d *Database) GetUserByUsername(username string) (*models.User, error) {
	query := `
		SELECT id, username, password, email
		FROM users
		WHERE username = $1
	`
	user := &models.User{}
	err := d.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error fetching user: %w", err)
	}
	return user, nil
}

func (d *Database) GetUserByID(id int) (*models.User, error){
	query := `
		SELECT id, username, password, email
		FROM users
		WHERE username = $1
	`
	user := &models.User{}
	err := d.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("erro fetching user: %w", err)
	}
	return user, nil
}