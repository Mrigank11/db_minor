package views

import (
	//"github.com/Mrigank11/db_minor/db"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//requires login
func AddToCart(w http.ResponseWriter, r *http.Request, username string) {
	params := mux.Vars(r)
	item_id := params["item_id"]
	fmt.Println(item_id)
}
