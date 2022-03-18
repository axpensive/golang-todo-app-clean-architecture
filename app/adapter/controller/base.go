package controller

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/view/%s.html", file))
	}
	// // エラーの記述省略でmust, エラーが起きたらpanicになる。
	templates := template.Must(template.ParseFiles(files...))
	log.Print(templates)
	templates.ExecuteTemplate(w, "layout", data)
}

// func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
// 	cookie, err := r.Cookie("_cookie")
// 	if err == nil {
// 		sess = models.Session{UUID: cookie.Value}
// 		if ok, _ := sess.CheckSession(); !ok {
// 			err = fmt.Errorf("invalid session")
// 		}
// 	}
// 	return sess, err
// }
