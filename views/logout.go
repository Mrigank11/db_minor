package views

import (
	//"github.com/Mrigank11/db_minor/db"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	delete(session.Values, "username")
	session.AddFlash("Logout Success")
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}
