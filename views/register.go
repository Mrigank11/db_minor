package views

import (
	"github.com/Mrigank11/db_minor/db"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "register", nil)
}

func RegisterHandlePost(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	//login user
	username := r.FormValue("username")
	password := r.FormValue("password")
	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")
	email := r.FormValue("email")
	mobile := r.FormValue("mobile")

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	password = string(bytes)
	_, err = db.DB.Query("insert into users values(?, ?, ?, ?, ?, NULL, ?, 2)", username, first_name, last_name, mobile, email, password)
	if err != nil {
		log.Debug(err)
		session.AddFlash("Username exists")
		session.Save(r, w)
		renderTemplate(w, r, "register", nil)
	} else {
		_ = db.Query("insert into cart(user_id) values(?)", username)
		session.Values["username"] = username
		session.AddFlash("Register success")
		session.Save(r, w)
		http.Redirect(w, r, "/", 302)
	}
}
