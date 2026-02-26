package routes

import (
		"BlogApp/handlers"
		"github.com/gorilla/mux"
		"net/http"
		"BlogApp/middleware"
)	

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/post", middleware.ValidateToken(http.HandlerFunc(handlers.CreatePost))).Methods("POST")
	r.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
	r.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/login", handlers.LoginUser).Methods("POST")
	r.HandleFunc("/users/{username}", handlers.GetUserByUsername).Methods("GET")
	r.Handle("/comments", middleware.ValidateToken(http.HandlerFunc(handlers.CreateComment))).Methods("POST")
	r.HandleFunc("/posts/{id}/comments", handlers.GetCommentsByPostID).Methods("GET")
	r.HandleFunc("/post/{id}", handlers.GetPostByID).Methods("GET")
	r.HandleFunc("/post/{id}", handlers.UpdatePostByID).Methods("PUT")
	r.HandleFunc("/post/{id}", handlers.DeletePostByID).Methods("DELETE")
	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	
	return r

}