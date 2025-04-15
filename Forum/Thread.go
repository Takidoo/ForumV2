package Forum

import (
	"Forum/Database"
	"database/sql"
)

type Thread struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Replies  int    `json:"replies"`
	Creation string `json:"creation"`
}

func GetLastedThreads(limit int) []Thread {
	rows, err := Database.DB.Query("SELECT id,title,created_at,replies FROM threads WHERE visible=true ORDER BY id DESC LIMIT ?;", limit)
	var Threads []Thread
	if err != nil {
		print(err.Error())
		return []Thread{}
	}
	for rows.Next() {
		var thread Thread
		rows.Scan(&thread.ID, &thread.Title, &thread.Creation, &thread.Replies)
		Threads = append(Threads, thread)
	}
	return Threads
}

func GetMostLikedThreads(limit int) []Thread {
	rows, err := Database.DB.Query("SELECT id,title,created_at,replies,user_id FROM threads WHERE visible=true ORDER BY likes ASC LIMIT ?;", limit)
	var Threads []Thread
	if err != nil {
		print(err.Error())
		return []Thread{}
	}
	for rows.Next() {
		var thread Thread
		var user_id int
		rows.Scan(&thread.ID, &thread.Title, &thread.Creation, &thread.Replies, &user_id)
		thread.Author = GetUserById(user_id).Username
		Threads = append(Threads, thread)
	}
	return Threads
}

func CheckIfThreadExist(thread_id string) bool {
	var title string
	query := `SELECT title FROM threads WHERE id = ?`
	err := Database.DB.QueryRow(query, thread_id).Scan(&title)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}
	return true
}
