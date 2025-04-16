package Forum

import (
	"Forum/Database"
	"fmt"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
}

func GetUser(session string) (User, error) {
	var user User
	err := Database.DB.QueryRow("SELECT user_id FROM sessions WHERE token = ?", session).Scan(&user.ID)
	if err != nil {

		return User{}, fmt.Errorf("cannot fetch user infos")
	}
	_ = Database.DB.QueryRow("SELECT username, role FROM users WHERE id = ?", user.ID).Scan(&user.Username, &user.Role)
	return user, nil
}
func GetUserById(user_id int) User {
	var user User
	err := Database.DB.QueryRow("SELECT username, role FROM users WHERE id = ?", user_id).Scan(&user.Username, &user.Role)
	if err != nil {
		return User{}
	}
	return user
}

func UserIsLogged(w http.ResponseWriter, r *http.Request) bool {
	session, err := r.Cookie("session_id")
	if err != nil {
		return false
	}
	var user_id int
	query := `SELECT user_id FROM sessions WHERE token = ?`
	err = Database.DB.QueryRow(query, session.Value).Scan(&user_id)
	if err != nil {
		return false
	}
	return true
}
func UserIsAdmin(session string) bool {
	var userID int
	err := Database.DB.QueryRow("SELECT user_id FROM sessions WHERE token=?", session).Scan(&userID)
	if err != nil {
		return false
	}
	var userRole int
	err = Database.DB.QueryRow("SELECT role FROM users WHERE id=?", userID).Scan(&userRole)
	if err != nil {
		return false
	}
	if userRole != 2 {
		return false
	}
	return true

}
