package Forum

import "Forum/Database"

type Post struct {
	PostID   int    `json:"post_id"`
	ThreadID int    `json:"thread_id"`
	Owner    int    `json:"owner"`
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
	rows, err := Database.DB.Query("SELECT id,user_id,content,created_at", thread_id)
	var posts []Post
	if err != nil {
		return posts
	}

	for rows.Next() {
		var post Post
		rows.Scan(&post.PostID, &post.Owner, &post.Content, &post.Date)
		post.Date = TimeAgo(post.Date)
		posts = append(posts, post)
	}
	return posts
}
