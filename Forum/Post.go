package Forum

import "Forum/Database"

type Post struct {
	PostID   int    `json:"post_id"`
	ThreadID int    `json:"thread_id"`
	Owner    string `json:"owner"`
	Content  string `json:"content"`
	Date     string `json:"created_at"`
}

func CreatePost(thread_id int, user_id int, content string) bool {
	_, err := Database.DB.Exec("INSERT INTO posts (thread_id,user_id,content) VALUES (?,?,?)", thread_id, user_id, content)
	if err != nil {
		return false
	}
	Database.DB.Exec("UPDATE threads SET replies =replies+1 WHERE id=?", thread_id)
	return true
}

func GetThreadPosts(thread_id string) []Post {
	rows, err := Database.DB.Query("SELECT id,user_id,content,created_at FROM posts WHERE thread_id=?", thread_id)
	var posts []Post
	if err != nil {
		return posts
	}

	for rows.Next() {
		var post Post
		var userID int
		rows.Scan(&post.PostID, &userID, &post.Content, &post.Date)
		post.Owner = GetUserById(userID).Username
		post.Date = TimeAgo(post.Date)
		posts = append(posts, post)
	}
	return posts
}
