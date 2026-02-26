package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"BlogApp/models"
	"BlogApp/tables"
	"github.com/gorilla/mux"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	email := r.Context().Value("email").(string)
	username := r.Context().Value("username").(string)

	comment.Email = email
	comment.Author = username

	id, err := tables.InsertComment(comment)
	if err != nil {
		http.Error(w, "Failed to create comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Comment created successfully", "id": id})
}

func GetCommentsByPostID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	params := mux.Vars(r)
	idParam := params["id"]

	postID, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "Invalid post_id parameter", http.StatusBadRequest)
		return
	}

	comments, err := tables.GetCommentsByPostID(postID)
	if err != nil {
		http.Error(w, "Failed to fetch comments", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}