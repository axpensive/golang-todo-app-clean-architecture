package controller

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/axpensive/golang-todo-app-clean-architecture/app/adapter/gateway/repository"
	"github.com/axpensive/golang-todo-app-clean-architecture/app/usecase/port"
)

type User struct {
	OutputFactory     func(w http.ResponseWriter) port.IUserOutputPort
	InputFactory      func(o port.IUserOutputPort, r port.IUserRepository) port.IUserInputPort
	RepositoryFactory func(c *sql.DB) port.IUserRepository
	Conn              *sql.DB
}

func Signup(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/login", 302)
		}
	case "POST":
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		user := repository.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/", 302)
	default:
		w.WriteHeader(405)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "login")
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	user, err := repository.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", 302)
	}

	if user.Password == repository.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}

		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}

	if err != http.ErrNoCookie {
		session := repository.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}
	http.Redirect(w, r, "/login", 302)
}
