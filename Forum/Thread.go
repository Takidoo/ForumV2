package Forum

import (
	"Forum/Database"
	"database/sql"
	"fmt"
	"time"
)

type Thread struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"creation"`
	Likes    int    `json:"likes"`
	Replies  int    `json:"replies"`
	Creation string `json:"creation"`
	Category int    `json:"category"`
}

func GetLastedThreads(limit int) []Thread {
	rows, err := Database.DB.Query("SELECT id,title,created_at,replies,user_id FROM threads WHERE visible=true ORDER BY id DESC LIMIT ?;", limit)
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
		thread.Creation = TimeAgo(thread.Creation)
		Threads = append(Threads, thread)
	}
	return Threads
}

func GetMostLikedThreads(limit int) []Thread {
	rows, err := Database.DB.Query("SELECT id,title,created_at,replies,user_id FROM threads WHERE visible=true ORDER BY replies ASC LIMIT ?;", limit)
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
		thread.Creation = TimeAgo(thread.Creation)
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

func TimeAgo(dateStr string) string {
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return "date invalide"
	}

	duration := time.Since(t)

	if duration < time.Minute {
		return "il y a quelques secondes"
	} else if duration < time.Hour {
		minutes := int(duration.Minutes())
		if minutes == 1 {
			return "il y a 1 minute"
		}
		return fmt.Sprintf("il y a %d minutes", minutes)
	} else if duration < 24*time.Hour {
		heures := int(duration.Hours())
		if heures == 1 {
			return "il y a 1 heure"
		}
		return fmt.Sprintf("il y a %d heures", heures)
	} else if duration < 30*24*time.Hour {
		jours := int(duration.Hours() / 24)
		if jours == 1 {
			return "il y a 1 jour"
		}
		return fmt.Sprintf("il y a %d jours", jours)
	} else if duration < 12*30*24*time.Hour {
		mois := int(duration.Hours() / (24 * 30))
		if mois == 1 {
			return "il y a 1 mois"
		}
		return fmt.Sprintf("il y a %d mois", mois)
	} else {
		annees := int(duration.Hours() / (24 * 365))
		if annees == 1 {
			return "il y a 1 an"
		}
		return fmt.Sprintf("il y a %d ans", annees)
	}
}
