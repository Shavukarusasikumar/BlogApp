package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	conString := "postgresql://neondb_owner:npg_PK7FYpVuXO8s@ep-young-art-aik4okaf-pooler.c-4.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require"

	db, err := sql.Open("postgres", conString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	
	DB = db
}