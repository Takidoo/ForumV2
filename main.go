package main

import (
	"Forum/API"
	"Forum/Database"
	"Forum/PageHandlers"
	"net/http"
)

func main() {
	// Prérequis
	Database.ConnectDB()
	Database.DB.Exec("INSERT INTO users (id, username, password, role) VALUES (1, 'Takido', '$2a$10$xGGhdn9iReF/EnZyP5iv9O9Rb3R2OWCsu/gLcWa849yclkvQFKqi.', 2)")
	http.Handle("/Resources/", http.StripPrefix("/Resources/", http.FileServer(http.Dir("./Resources"))))

	// API
	http.HandleFunc("/api/login", API.Login)
	http.HandleFunc("/api/register", API.Register)
	http.HandleFunc("/api/FetchThreadPosts", API.FetchThreadPosts)
	http.HandleFunc("/api/CreateThread", API.CreateThread)
	http.HandleFunc("/api/CreatePost", API.CreatePost)
	http.HandleFunc("/api/SetUserRole", API.SetUserRole)
	http.HandleFunc("/api/DisableAccount", API.DisableAccount)
	http.HandleFunc("/api/CreateCategory", API.CreateCategory)

	// Pages
	http.HandleFunc("/", PageHandlers.HomePageHandler)
	http.HandleFunc("/test", PageHandlers.TestPageHandler)
	http.HandleFunc("/admin", PageHandlers.AdminPageHandler)
	http.HandleFunc("/login", PageHandlers.LoginHandler)
	http.HandleFunc("/register", PageHandlers.RegisterHandler)

	// Démarage du serveur
	http.ListenAndServe(":80", nil)
}
