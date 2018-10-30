package views

import (
	//"github.com/Mrigank11/db_minor/db"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "register", nil)
}
