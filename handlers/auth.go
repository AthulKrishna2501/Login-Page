package auth

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

type loginErr struct {
	UserNameErr string
	PassWordErr string
}

var (
	Username = "athul"
	Password = "loginaccess"
	store    = sessions.NewCookieStore([]byte("super-key"))
)

func Login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	if session.Values["username"] != nil {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	var err loginErr
	var username string

	if r.Method == http.MethodPost {
		username = r.FormValue("username")
		password := r.FormValue("password")

		if username != Username {
			err.UserNameErr = "Invalid Username"
		}

		if password != Password {
			err.PassWordErr = "Invalid Password"
		}

		// Check if there are no errors
		if err.UserNameErr == "" && err.PassWordErr == "" {
			session.Values["username"] = username
			session.Save(r, w)
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	tmpl.Execute(w, err)
}

func Home(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	user := session.Values["username"]
	if user == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	http.ServeFile(w, r, "templates/home.html")
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
