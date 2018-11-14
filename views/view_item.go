package views

import (
	"net/http"

	"github.com/Mrigank11/db_minor/db"
	//log "github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
)

type tag struct {
	ID          int
	Name        string
	Description string
}

func ViewItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	item_id := params["item_id"]

	//get all tags
	rows := db.Query(`
		select t.* from product_tags pt 
		join products p on p.product_id = pt.product_id and p.product_id = ?
		join tag t on t.tag_id = pt.tag_id;
	`, item_id)
	var tags []tag
	//load products
	for rows.Next() {
		var p tag
		rows.Scan(&p.ID, &p.Name, &p.Description)
		tags = append(tags, p)
	}

	rows = db.Query(`select * from products_with_best_prices where product_id = ?;`, item_id)
	var product product
	//load products
	if rows.Next() {
		var sid int
		rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Dealer, &product.Remaining, &sid)
	}
	//finally render
	renderTemplate(w, r, "item", map[string]interface{}{
		"Product": product,
		"Tags":    tags,
	})
}
