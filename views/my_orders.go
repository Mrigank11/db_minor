package views

import (
	"net/http"

	"github.com/Mrigank11/db_minor/db"
	//log "github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
)

type sale struct {
	TID    string
	Amount float64
	Date   string
}

func MyOrders(w http.ResponseWriter, r *http.Request, username string) {
	rows := db.Query(`
		select s.* from sale s
		join cart c on s.transaction_id = c.transaction_id
		join users u on u.username = c.user_id and u.username = ? 
	`, username)

	var sales []sale
	//load sales
	for rows.Next() {
		var p sale
		rows.Scan(&p.TID, &p.Amount, &p.Date)
		sales = append(sales, p)
	}
	//finally render
	renderTemplate(w, r, "my_orders", sales)
}

func ViewOldCart(w http.ResponseWriter, r *http.Request, username string) {
	params := mux.Vars(r)
	tid := params["tid"]

	rows := db.Query(`select p.*, s.amount from sale s
	join cart c on s.transaction_id = c.transaction_id and c.transaction_id = ?
	join cart_items ci on ci.cart_id = c.cart_id
	join products p on p.product_id = ci.product_id;`, tid)

	var products []product
	var total float64
	//load products
	for rows.Next() {
		var p product
		rows.Scan(&p.ID, &p.Name, &p.Description, &total)
		products = append(products, p)
	}

	//finally render
	renderTemplate(w, r, "old_cart", map[string]interface{}{
		"Total":    total,
		"Products": products,
	})
}
