package API

import (
	"Forum/Database"
	"Forum/Forum"
	"net/http"
	"strconv"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	middleAuth, aerr := Database.MiddlewareAuth(w, r)
	if !middleAuth {
		http.Error(w, aerr.Error(), http.StatusUnauthorized)
		return
	}

	thread_id := r.FormValue("thread_id")
	message := r.FormValue("message")
	if message == "" || thread_id == "" {
		http.Error(w, "Invalid Args", http.StatusBadRequest)
		return
	}

	if !Forum.CheckIfThreadExist(thread_id) {
		http.Error(w, "Thread doesn't exist", http.StatusBadRequest)
		return
	}

	cookie, _ := r.Cookie("session_id")
	user, err := Forum.GetUser(cookie.Value)
	if err != nil {
		http.Error(w, "Cannot get user info", http.StatusInternalServerError)
		return
	}
	tid, _ := strconv.Atoi(thread_id)
	Forum.CreatePost(tid, user.ID, message)
}
