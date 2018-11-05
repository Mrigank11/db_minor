package views

import (
	"fmt"
	"github.com/Mrigank11/db_minor/db"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//requires login
func AddToCart(w http.ResponseWriter, r *http.Request, username string) {
	session, _ := store.Get(r, "session")
	params := mux.Vars(r)
	item_id := params["item_id"]
	cart_id := getCartId(r)
	db.Query("insert into cart_items(product_id, cart_id) values(?, ?)", item_id, cart_id)
	session.AddFlash(fmt.Sprintf("Item %s added", item_id))
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}

func RmFromCart(w http.ResponseWriter, r *http.Request, username string) {
	session, _ := store.Get(r, "session")
	params := mux.Vars(r)
	item_id := params["item_id"]
	cart_id := getCartId(r)
	db.Query("delete from cart_items where cart_id = ? and product_id = ?", cart_id, item_id)
	session.AddFlash("Item deleted")
	session.Save(r, w)

	http.Redirect(w, r, "/cart", 302)
}

func CheckoutCart(w http.ResponseWriter, r *http.Request, username string) {
	session, _ := store.Get(r, "session")
	cart_id := getCartId(r)

	var amount float64
	row := db.Query("select sum(p.price) from cart_items c inner join products_with_best_prices p on p.product_id = c.product_id and cart_id=?", cart_id)
	if row.Next() {
		row.Scan(&amount)
	}

	//ideally, this will be sent by the payment gateway
	transaction_id := randSeq(10)
	log.Info("transaction created ", transaction_id)
	db.Query("insert into sale(transaction_id, amount) values(?, ?)", transaction_id, amount)
	db.Query("update cart set transaction_id=? where cart_id=?", transaction_id, cart_id)

	//TODO: do this with triggers
	_ = db.Query("insert into cart(user_id) values(?)", username)

	session.AddFlash("Cart checked out successfully")
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}
