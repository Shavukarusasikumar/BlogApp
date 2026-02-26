package main

import (
	"BlogApp/config"
	"BlogApp/routes"
	"BlogApp/tables"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	config.ConnectDB()
	fmt.Println("Database connection established successfully!")

	tables.CreateUserTable()
	tables.CreatePostTable()
	tables.CreateCommentTable()

	router := routes.SetupRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port:", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}