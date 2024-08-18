package auth

import (
	"encoding/json"
	"errors"
	"github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	db "registerlogin/database"
	"registerlogin/models"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside RegisterHandler")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Decoded User: %+v", user)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	_, err = db.DB.Exec("INSERT INTO user (username, password) VALUES (?, ?)",
		user.Username, string(hashedPassword))
	if err != nil {
		var sqlErr sqlite3.Error
		if errors.As(err, &sqlErr) && errors.Is(sqlErr.Code, sqlite3.ErrConstraint) {
			http.Error(w, "Username already taken", http.StatusConflict)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message":  "Signed up successfully",
		"username": user.Username,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	log.Printf("Signed up as %s", user.Username)
	w.WriteHeader(http.StatusCreated)
}
