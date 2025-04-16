package Forum

import (
	"Forum/Database"
	"database/sql"
)

func Search(input string) []Thread {
	print("Searching...")
	var result []Thread
	if input != "" {
		print("Search by name")
		SearchByName(input, &result)
	}
	return result
}

func SearchByName(input string, threads *[]Thread) {
	print(input)
	rows, err := Database.DB.Query("SELECT id, user_id, title, created_at, category, likes, replies FROM threads WHERE title LIKE ? AND visible=true", "%"+input+"%")
	if err != nil {
		print(err.Error())
		if err == sql.ErrNoRows {
			print("recherche de merde")
		}
		return
	}
	print("test")
	defer rows.Close()

	for rows.Next() {
		print("cool")
		var thread Thread
		var userID int
		var time string
		rows.Scan(&thread.ID, &userID, &thread.Title, &time, &thread.Category, &thread.Likes, &thread.Replies)
		thread.Author = GetUserById(userID).Username
		thread.Creation = TimeAgo(time)
		*threads = append(*threads, thread)
	}
	print("test2")
}
