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
	Price       float64
	Dealer      string
	Remaining   int
}

func Index(w http.ResponseWriter, r *http.Request) {
	rows := db.Query(`select * from products_with_best_prices;`)
	products := getProductList(rows)
	//finally render
	renderTemplate(w, r, "home", products)
}
