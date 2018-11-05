package views

import (
	"net/http"

	"github.com/Mrigank11/db_minor/db"
	//log "github.com/sirupsen/logrus"
)

func Cart(w http.ResponseWriter, r *http.Request) {
	cart_id := getCartId(r)

	rows := db.Query("select p.* from cart_items c inner join products_with_best_prices p on p.product_id = c.product_id and cart_id=?", cart_id)
	var products []product
	//load products
	for rows.Next() {
		var p product
		rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Dealer, &p.Remaining)
		products = append(products, p)
	}

	total := 0.0
	for i := range products {
		total += products[i].Price
	}
	//finally render
	renderTemplate(w, r, "cart", map[string]interface{}{
		"Total":    total,
		"Products": products,
	})
}
