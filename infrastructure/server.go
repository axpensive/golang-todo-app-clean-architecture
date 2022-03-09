package infrastructure

import (
	"io"
	"net/http"
)

func StartAppServer() error {
	test := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Test")
	}
	http.HandleFunc("/", test)
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
	return http.ListenAndServe("localhost:7777", nil)
}
