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

type SearchPageData struct {
	Result   []Forum.Thread
	IsLogged bool
	IsAdmin  bool
}

type ThreadPageData struct {
	Posts    []Forum.Post
	IsLogged bool
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
	if Forum.UserIsLogged(w, r) {
		tmpl.Execute(w, HomePageData{
			Username:        "",
			LastedThreads:   Forum.GetLastedThreads(10),
			MostLikedThread: Forum.GetMostLikedThreads(10),
			IsLogged:        true,
			IsAdmin:         false,
		})
		return
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
func CreationHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("WebPages/creation.html")
	if err != nil {
		log.Println("Erreur lors du chargement de la page de création :", err)
		return
	}
	tmpl.Execute(w, nil)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("WebPages/search.html")
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		if Forum.UserIsLogged(w, r) {
			tmpl.Execute(w, SearchPageData{
				Result:   Forum.Search(r.FormValue("input")),
				IsLogged: true,
				IsAdmin:  false,
			})
			return
		}
		tmpl.Execute(w, SearchPageData{
			Result:   Forum.Search(r.FormValue("input")),
			IsLogged: false,
			IsAdmin:  false,
		})
		return
	}

	if Forum.UserIsLogged(w, r) {
		tmpl.Execute(w, SearchPageData{
			Result:   Forum.Search(r.FormValue("input")),
			IsLogged: true,
			IsAdmin:  false,
		})
		return
	}
	tmpl.Execute(w, SearchPageData{
		Result:   Forum.Search(r.FormValue("input")),
		IsLogged: false,
		IsAdmin:  false,
	})

}

func ThreadHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("WebPages/thread.html")
	if !Forum.CheckIfThreadExist(r.URL.Query().Get("thread_id")) {
		http.Error(w, "Invalid Thread ID", http.StatusBadRequest)
		return
	}
	if Forum.UserIsLogged(w, r) {
		tmpl.Execute(w, ThreadPageData{
			Posts:    Forum.GetThreadPosts(r.URL.Query().Get("thread_id")),
			IsLogged: true,
		})
		return
	}
	tmpl.Execute(w, ThreadPageData{
		Posts:    Forum.GetThreadPosts(r.URL.Query().Get("thread_id")),
		IsLogged: false,
	})
}
