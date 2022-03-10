package infrastructure

import (
	"log"
	"net/http"
)

func StartAppServer() {
	// http.HandleFunc("/", top)
	// http.HandleFunc("/signup", signup)
	// http.HandleFunc("/login", login)
	// http.HandleFunc("/authenticate", authenticate)
	// http.HandleFunc("/logout", logout)
	// http.HandleFunc("/todos", index)
	// http.HandleFunc("/todos/new", todoNew)
	// http.HandleFunc("/todos/save", todoSave)
	// // edit/1のようなURLを判定したので/をつける
	// http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	// http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	// http.HandleFunc("/todos/delete/", parseURL(todoDelete))
	err := http.ListenAndServe("localhost:7777", nil)
	if err != nil {
		log.Fatalf("Server down: %s", err)
	}
}
