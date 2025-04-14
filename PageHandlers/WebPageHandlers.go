package PageHandlers

import (
	"Forum/Forum"
	"html/template"
	"log"
	"net/http"
)

type HomePageData struct {
	Username        string
	LastedThreads   []Forum.Thread
	MostLikedThread []Forum.Thread
	IsLogged        bool
	IsAdmin         bool
}

type AdminPageData struct {
	Username string
}

type LoginRegisterData struct {
	Type string
}

func TestPageHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl, _ = template.ParseFiles("WebPages/test.html")
	tmpl.Execute(w, nil)
}
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl, _ = template.ParseFiles("WebPages/forum.html")
	cookie, err := r.Cookie("session_id")
	if err == nil {
		user, err := Forum.GetUser(cookie.Value)
		if err == nil {
			tmpl.Execute(w, HomePageData{
				Username:        user.Username,
				LastedThreads:   Forum.GetLastedThreads(10),
				MostLikedThread: Forum.GetMostLikedThreads(10),
				IsLogged:        true,
				IsAdmin:         user.Role == 2,
			})
			return
		}
	}

	tmpl.Execute(w, HomePageData{
		Username:        "",
		LastedThreads:   Forum.GetLastedThreads(10),
		MostLikedThread: Forum.GetMostLikedThreads(10),
		IsLogged:        false,
		IsAdmin:         false,
	})
}
func AdminPageHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := r.Cookie("session_id")
	if !Forum.UserIsAdmin(session.Value) {
		http.Error(w, "You're not admin", http.StatusUnauthorized)
		return
	}
	cookie, _ := r.Cookie("session_id")
	user, _ := Forum.GetUser(cookie.Value)
	var tmpl, _ = template.ParseFiles("WebPages/admin.html")
	tmpl.Execute(w, AdminPageData{
		Username: user.Username,
	})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("WebPages/login.html")
	var PageData LoginRegisterData = LoginRegisterData{"wrapper"}
	if err != nil {
		log.Println("Erreur lors du chargement de la page de connexion :", err)
		return
	}
	tmpl.Execute(w, PageData)
}
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("WebPages/login.html")
	var PageData LoginRegisterData = LoginRegisterData{"wrapper active"}
	if err != nil {
		log.Println("Erreur lors du chargement de la page de connexion :", err)
		return
	}
	tmpl.Execute(w, PageData)
}
