package auth

import (
	"database/sql"
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	db "registerlogin/database"
	"registerlogin/models"
)

func DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside DeleteAccountHandler")
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if request.Username == "" || request.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	var user models.User
	err = db.DB.QueryRow("SELECT id, username, password FROM user WHERE username = ?",
		request.Username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	_, err = db.DB.Exec("DELETE FROM user WHERE username = ?", request.Username)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message":  "Successfully deleted",
		"username": user.Username,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	log.Printf("Account deleted for user %s", request.Username)
}
