package views

import (
	"net/http"

	"github.com/Mrigank11/db_minor/db"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "login", nil)
}

func LoginHandlePost(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	login_success := false
	//login user
	username := r.FormValue("username")
	password := r.FormValue("password")
	rows := db.Query("select password from users where username = ? limit 1", username)
	if rows.Next() {
		var hash []byte
		rows.Scan(&hash)
		//check pass
		err := bcrypt.CompareHashAndPassword(hash, []byte(password))
		if err == nil {
			login_success = true
		}
	}
	if login_success {
		log.Debug(username, " logged in")
		session.Values["username"] = username
		//  TODO: any better way? <23-10-18, Mrigank11> //
		session.Save(r, w)
		http.Redirect(w, r, "/", 302)
	} else {
		session.AddFlash("Login Failed")
	}
	session.Save(r, w)
	renderTemplate(w, r, "login", nil)
}
