package tables

import (
	"fmt"
	"BlogApp/config"
	"log"
)

func CreateUserTable() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) NOT NULL UNIQUE,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := config.DB.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Users table created successfully!")
}

func RegisterUser(Username, Email, Password string) error {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
	_, err := config.DB.Exec(query, Username, Email, Password)
	return err
}