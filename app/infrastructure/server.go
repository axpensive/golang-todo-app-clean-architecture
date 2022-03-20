package infrastructure

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/axpensive/golang-todo-app-clean-architecture/app/adapter/controller"
)

var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// URLがマッチするか
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}

		// urlにIDが入っているか。
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, qi)
	}
}

func StartAppServer() {
	http.HandleFunc("/", controller.TopPage)
	http.HandleFunc("/signup", controller.Signup)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/authenticate", controller.Authenticate)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/todos", controller.TodoList)
	http.HandleFunc("/todos/new", controller.CreateTodo)
	http.HandleFunc("/todos/save", controller.SaveTodo)
	// edit/1のようなURLを判定したので/をつける
	http.HandleFunc("/todos/edit/", parseURL(controller.EditTodo))
	http.HandleFunc("/todos/update/", parseURL(controller.UpdateTodo))
	http.HandleFunc("/todos/delete/", parseURL(controller.DeleteTodo))
	err := http.ListenAndServe("localhost:7777", nil)
	if err != nil {
		log.Fatalf("Server down: %s", err)
	}
}
