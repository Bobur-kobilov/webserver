package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/webserver/utils"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func CheckToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := utils.CheckToken(r)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(token)
		next.ServeHTTP(w, r)
	})
}
