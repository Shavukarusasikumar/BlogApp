package main

import (
	"BlogApp/config"
	"BlogApp/routes"
	"BlogApp/tables"
	"fmt"
	"net/http"
	"log"
)

func main() {
	config.ConnectDB()
	fmt.Println("Database connection established successfully!")

	tables.CreateUserTable()
	tables.CreatePostTable()
	tables.CreateCommentTable()

	routes.SetupRoutes()
	fmt.Println("Server is running on http://localhost:8080")
	
	router := routes.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8080", router))
}

