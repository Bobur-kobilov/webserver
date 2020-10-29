package router

import (
	"log"
	"net/http"

	"database/sql"

	"github.com/gorilla/mux"
	"github.com/webserver/handler"
)

func Router(DB *sql.DB) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/signup", handler.SignUp(DB)).Methods("POST")
	router.HandleFunc("/login", handler.Login(DB)).Methods("POST")
	router.HandleFunc("/data", handler.RegisterData(DB)).Methods("POST")
	router.HandleFunc("/data", handler.QueryData(DB))

	log.Fatal(http.ListenAndServe(":10000", router))
}