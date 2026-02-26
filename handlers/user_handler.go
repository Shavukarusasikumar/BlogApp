package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	"BlogApp/config"
	"BlogApp/models"
	"BlogApp/utils"
	"BlogApp/tables"
)

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	
	err = tables.RegisterUser(user.Username, user.Email, hashedPassword)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var user models.User
	var storedHashedPassword string

	query := `SELECT username, password FROM users WHERE email = $1`
	err = config.DB.QueryRow(query, credentials.Email).Scan(&user.Username, &storedHashedPassword)
	if err != nil {
		http.Error(w, "Email not exists", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(credentials.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(credentials.Email, user.Username)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	params := mux.Vars(r)

	userName := params["username"]

	if userName == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	query := `SELECT id, username, email, created_at FROM users WHERE username = $1`

	var user models.User

	err := config.DB.QueryRow(query, userName).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
