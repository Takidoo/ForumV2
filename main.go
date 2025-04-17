package main

import (
	"Forum/API"
	"Forum/Database"
	"Forum/PageHandlers"
	"net/http"
)

func InsertDbDefalut() {
	Database.DB.Exec("INSERT INTO users (id, username, password, role) VALUES (1, 'Takido', '$2a$10$xGGhdn9iReF/EnZyP5iv9O9Rb3R2OWCsu/gLcWa849yclkvQFKqi.', 2)")
	Database.DB.Exec("INSERT INTO categories (name) VALUES ('Espagnol')")
	Database.DB.Exec("INSERT INTO categories (name) VALUES ('Français')")
	Database.DB.Exec("INSERT INTO categories (name) VALUES ('Italien')")
	Database.DB.Exec("INSERT INTO categories (name) VALUES ('Portugais')")
	Database.DB.Exec("INSERT INTO categories (name) VALUES ('Thai')")
	Database.DB.Exec("INSERT INTO categories (name) VALUES ('English')")
}

func main() {
	// Prérequis
	Database.ConnectDB()
	http.Handle("/Resources/", http.StripPrefix("/Resources/", http.FileServer(http.Dir("./Resources"))))
	InsertDbDefalut()

	// API
	http.HandleFunc("/api/login", API.Login)
	http.HandleFunc("/api/register", API.Register)
	http.HandleFunc("/api/FetchThreadPosts", API.FetchThreadPosts)
	http.HandleFunc("/api/CreateThread", API.CreateThread)
	http.HandleFunc("/api/CreatePost", API.CreatePost)
	http.HandleFunc("/api/SetUserRole", API.SetUserRole)
	http.HandleFunc("/api/DisableAccount", API.DisableAccount)
	http.HandleFunc("/api/CreateCategory", API.CreateCategory)
	http.HandleFunc("/api/LikeThread", API.LikeThread)

	// Pages
	http.HandleFunc("/", PageHandlers.HomePageHandler)
	http.HandleFunc("/test", PageHandlers.TestPageHandler)
	http.HandleFunc("/admin", PageHandlers.AdminPageHandler)
	http.HandleFunc("/login", PageHandlers.LoginHandler)
	http.HandleFunc("/register", PageHandlers.RegisterHandler)
	http.HandleFunc("/creation", PageHandlers.CreationHandler)
	http.HandleFunc("/search", PageHandlers.SearchHandler)
	http.HandleFunc("/thread", PageHandlers.ThreadHandler)

	// Démarage du serveur
	http.ListenAndServe(":80", nil)
}
