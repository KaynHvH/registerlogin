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

func ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside ChangePasswordHandler")

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Username    string `json:"username"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if request.Username == "" || request.OldPassword == "" || request.NewPassword == "" {
		http.Error(w, "Username, old password, and new password are required", http.StatusBadRequest)
		return
	}

	var user models.User
	err = db.DB.QueryRow("SELECT id, username, password FROM user WHERE username = ?", request.Username).
		Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.OldPassword))
	if err != nil {
		http.Error(w, "Old password does not match", http.StatusUnauthorized)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	_, err = db.DB.Exec("UPDATE user SET password = ? WHERE username = ?", string(hashedPassword), request.Username)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "Password updated",
		"user":    user.Username,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}

	log.Printf("Changed password for %s account\n", request.Username)
	w.WriteHeader(http.StatusOK)
}
