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

///checkout cart
///this essentially means setting the transaction_id
///to the active user cart.
///MySQL triggers will handle the rest.
func CheckoutCart(w http.ResponseWriter, r *http.Request, username string) {
	session, _ := store.Get(r, "session")
	cart_id := getCartId(r)

	//ideally, this will be sent by the payment gateway
	transaction_id := randSeq(10)
	_, err := db.DB.Query("update cart set transaction_id=? where cart_id=?", transaction_id, cart_id)
	if err != nil {
		log.Error(err)
		session.AddFlash(err.Error())
	} else {
		_ = db.Query("insert into cart(user_id) values(?)", username)
		log.Info("transaction created ", transaction_id)
		session.AddFlash("Cart checked out successfully")
	}

	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}
