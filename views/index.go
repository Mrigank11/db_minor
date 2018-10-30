package views

import (
	"net/http"

	"github.com/Mrigank11/db_minor/db"
	//log "github.com/sirupsen/logrus"
)

type product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	Dealer      string
	Remaining   int
}

func Index(w http.ResponseWriter, r *http.Request) {
	rows := db.Query(`
			select p.*, s.price, s.dealer, count(s1.price) from products p left join stock s on s.stock_id = (
				select stock_id from stock s_
					where s_.product_id = p.product_id
					order by s_.price asc
					limit 1
			) 
			left join stock s1 on s1.product_id = p.product_id
			group by p.product_id;
		`)
	var products []product
	//load products
	for rows.Next() {
		var p product
		rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Dealer, &p.Remaining)
		products = append(products, p)
	}
	//finally render
	renderTemplate(w, r, "home", products)
}
