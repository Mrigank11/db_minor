package views

import (
	"fmt"
	"github.com/Mrigank11/db_minor/db"
	"github.com/gorilla/mux"
	"net/http"
)

//requires login
func AddToCart(w http.ResponseWriter, r *http.Request, username string) {
	session, _ := store.Get(r, "session")
	params := mux.Vars(r)
	item_id := params["item_id"]
	rows := db.Query("select cart_id from cart where user_id = ? and transaction_id is NULL", session.Values["username"])
	rows.Next()
	var cart_id int
	rows.Scan(&cart_id)

	db.Query("insert into cart_items(product_id, cart_id) values(?, ?)", item_id, cart_id)
	session.AddFlash(fmt.Sprintf("Item %s added", item_id))
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}
func RmFromCart(w http.ResponseWriter, r *http.Request, username string) {
	session, _ := store.Get(r, "session")
	params := mux.Vars(r)
	item_id := params["item_id"]
	rows := db.Query("select cart_id from cart where user_id = ? and transaction_id is NULL", session.Values["username"])
	rows.Next()
	var cart_id int
	rows.Scan(&cart_id)

	db.Query("delete from cart_items where cart_id = ? and product_id = ?", cart_id, item_id)
	session.AddFlash("Item deleted")
	session.Save(r, w)

	http.Redirect(w, r, "/cart", 302)
}
