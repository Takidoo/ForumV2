package API

import (
	"Forum/Database"
	"database/sql"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
	_, err := Database.MiddlewareAuth(w, r)
	if err == nil {
		http.Error(w, "Utilisateur déjà connecté", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}

	success, err := LoginUser(username, password, w)
	if err != nil {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}

	if !success {
		http.Error(w, "Invalid ids", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func LoginUser(username, password string, w http.ResponseWriter) (bool, error) {
	var storedPassword, token string
	var userID int
	var account_disabled bool

	query := "SELECT id, password, account_disabled FROM users WHERE username = ?"
	err := Database.DB.QueryRow(query, username).Scan(&userID, &storedPassword, &account_disabled)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	if account_disabled {
		return false, fmt.Errorf("Account is disabled, please contact support")
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		return false, fmt.Errorf("Can't hash password")
	}

	token = Database.GenerateToken()
	query = "INSERT INTO sessions (token, user_id) VALUES (?,?)"
	_, err = Database.DB.Exec(query, token, userID)
	if err != nil {
		return false, fmt.Errorf("Can't create session")
	}
	cookie := http.Cookie{
		Name:     "session_id",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	return true, nil
}
