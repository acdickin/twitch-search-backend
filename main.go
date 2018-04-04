package main

import (
"log"
"net/http"

"github.com/gorilla/mux"
"github.com/rs/cors"

)

func main() {

	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
	})
	
	
	router.HandleFunc("/stream/top", StreamsByIdFunc)
	router.HandleFunc("/user/{userid}", UserByIdFunc)
	router.HandleFunc("/game/{gameid}", GameByIdFunc)
	router.HandleFunc("/video/{videoid}", UserByIdFunc)
	
	handler := c.Handler(router)
	handler = Logger(handler, "String")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
