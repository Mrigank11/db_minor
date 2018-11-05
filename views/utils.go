package views

import (
	"fmt"
	"github.com/Mrigank11/db_minor/db"
	"github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"
	"html/template"
	"math/rand"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("super secret key"))

func renderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	session, _ := store.Get(r, "session")
	t, err := template.ParseFiles("./templates/layout/base.html", "./templates/layout/footer.html", "./templates/layout/nav.html", fmt.Sprintf("./templates/%s.html", name))
	if err != nil {
		log.Error(err)
	}
	flashes := session.Flashes()
	session.Save(r, w)
	t.Execute(w, map[string]interface{}{"Session": session.Values, "Flashes": flashes, "Data": data})
}

func RequiresLogin(handler func(http.ResponseWriter, *http.Request, string)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")
		if username, ok := session.Values["username"].(string); ok {
			handler(w, r, username)
		} else {
			//  TODO: add flash <24-10-18, Mrigank11> //
			http.Redirect(w, r, "/", 302)
		}
	}
}

func getCartId(r *http.Request) int {
	session, _ := store.Get(r, "session")
	rows := db.Query("select cart_id from cart where user_id = ? and transaction_id is NULL", session.Values["username"])
	rows.Next()
	var cart_id int
	rows.Scan(&cart_id)
	return cart_id
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
