package views

import (
	"net/http"

	"github.com/Mrigank11/db_minor/db"
	//log "github.com/sirupsen/logrus"
)

func Cart(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	rows := db.Query("select cart_id from cart where user_id = ? and transaction_id is NULL", session.Values["username"])
	rows.Next()
	var cart_id int
	rows.Scan(&cart_id)

	rows = db.Query("select product_id from cart_items where cart_id = ?", cart_id)
	var products []string
	//load products
	for rows.Next() {
		var item_id string
		rows.Scan(&item_id)
		products = append(products, item_id)
	}
	//finally render
	renderTemplate(w, r, "cart", products)
}
