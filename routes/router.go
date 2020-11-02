package router

import (
	"log"
	"net/http"

	"database/sql"

	"github.com/gorilla/mux"
	"github.com/webserver/handler"
	"github.com/webserver/middleware"
)

func Router(DB *sql.DB) {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.LoggingMiddleware)
	router.Use(mux.CORSMethodMiddleware(router))

	router.HandleFunc("/signup", handler.SignUp(DB)).Methods("POST")
	router.HandleFunc("/login", handler.Login(DB)).Methods("POST")

	router.Handle("/data", middleware.CheckToken(handler.RegisterData(DB))).Methods("POST")
	router.Handle("/data", middleware.CheckToken(handler.QueryData(DB)))

	log.Fatal(http.ListenAndServe(":10000", router))
}
