package views

import (
	"net/http"

	"github.com/Mrigank11/db_minor/db"
	//log "github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
)

func SearchByTag(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tag_id := params["tag_id"]

	//find tag details
	rows := db.Query("select * from tag where tag_id = ?", tag_id)
	var p tag
	if rows.Next() {
		rows.Scan(&p.ID, &p.Name, &p.Description)
	}

	rows = db.Query(`
		select ps.* from products_with_best_prices ps
		join product_tags p on p.product_id = ps.product_id
		join tag t on t.tag_id = p.tag_id and t.tag_id = ?;
	`, tag_id)

	products := getProductList(rows)
	//finally render
	renderTemplate(w, r, "tag_search", map[string]interface{}{
		"Tag":      p,
		"Products": products,
	})
}
