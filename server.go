package main

import (
	"flag"
	"fmt"
	"github.com/Mrigank11/db_minor/views"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	//parse flags
	port := flag.Int("port", 8080, "The port to listen on")
	flag.Parse()
	log.SetLevel(log.DebugLevel)

	//init
	router := mux.NewRouter()

	/*
	 *ROUTES
	 */
	router.HandleFunc("/", views.Index)
	router.HandleFunc("/login", views.Login).Methods("GET")
	router.HandleFunc("/login", views.LoginHandlePost).Methods("POST")
	router.HandleFunc("/register", views.Register).Methods("GET")
	router.HandleFunc("/register", views.RegisterHandlePost).Methods("POST")
	router.HandleFunc("/logout", views.Logout)

	router.HandleFunc("/cart", views.Cart)
	router.HandleFunc("/to_cart/{item_id}", views.RequiresLogin(views.AddToCart))
	router.HandleFunc("/rm_from_cart/{item_id}", views.RequiresLogin(views.RmFromCart))

	//start server
	log.Info("Listening on port: ", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}
