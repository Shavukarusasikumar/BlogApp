package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"BlogApp/config"
	"BlogApp/models"
	"BlogApp/tables"
	"strconv"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if(r.Method != http.MethodPost) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	email := r.Context().Value("email").(string)
	username := r.Context().Value("username").(string)

	query := `INSERT INTO posts (title, content, author, email, created_at, updated_at) VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id`
	err = config.DB.QueryRow(query, post.Title, post.Content, username, email).Scan(&post.ID)
	if err != nil {
		log.Printf("Error inserting post: %v", err)
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

type PaginatedResponse struct {
	Data        []models.Post `json:"data"`
	Total       int    `json:"total"`
	Page        int    `json:"page"`
	Limit       int    `json:"limit"`
	TotalPages  int    `json:"total_pages"`
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	if(r.Method != http.MethodGet) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	limit := 10
	page := 1
	sortBy := "created_at"
	sortOrder := "desc"

	queryParams := r.URL.Query()

	if l := queryParams.Get("limit"); l != "" {
		if v, ok := queryParams["limit"]; ok && len(v) > 0 {
			limit, _ = strconv.Atoi(v[0])
		}
	}

	if p := queryParams.Get("page"); p != "" {
		if v, ok := queryParams["page"]; ok && len(v) > 0 {
			page, _ = strconv.Atoi(v[0])
		}
	}

	//sorting
	if s := queryParams.Get("sort_by"); s != "" {
		if s == "created_at" || s == "title" || s == "author" {
			sortBy = s
		} else {
			http.Error(w, "Invalid sort_by parameter", http.StatusBadRequest)
			return
		}
	}

	if o := queryParams.Get("sort_order"); o != "" {
		if o == "asc" || o == "desc" {
			sortOrder = o
		} else {
			http.Error(w, "Invalid sort_order parameter", http.StatusBadRequest)
			return
		}
	}

	offset := (page - 1) * limit

	totalCount, err := tables.GetPostsCount()

	if err != nil {
		log.Printf("Error fetching posts count: %v", err)
		http.Error(w, "Failed to fetch posts count", http.StatusInternalServerError)
		return
	}

	posts, err := tables.GetAllPosts(limit, offset, sortBy, sortOrder)
	if err != nil {
		log.Printf("Error fetching posts: %v", err)
		http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}

	totalPages := (totalCount + limit - 1) / limit

	response := PaginatedResponse{
		Data:        posts,
		Total:       totalCount,
		Page:        page,
		Limit:       limit,
		TotalPages:  totalPages,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	if(r.Method != http.MethodGet) {
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

	post, err := tables.GetPostByID(postID)
	if err != nil {
		log.Printf("Error fetching post by ID: %v", err)
		http.Error(w, "Failed to fetch post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func UpdatePostByID(w http.ResponseWriter, r *http.Request) {
	if(r.Method != http.MethodPut) {
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

	var post models.Post
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	post.ID = postID

	err = tables.UpdatePost(post)
	if err != nil {
		log.Printf("Error updating post: %v", err)
		http.Error(w, "Failed to update post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Post updated successfully"})
}

func DeletePostByID(w http.ResponseWriter, r *http.Request) {
	if(r.Method != http.MethodDelete) {
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

	err = tables.DeletePost(postID)
	if err != nil {
		log.Printf("Error deleting post: %v", err)
		http.Error(w, "Failed to delete post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Post deleted successfully"})
}