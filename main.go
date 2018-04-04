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
	
	
	router.HandleFunc("/stream/top", StreamsByIdFunc).Methods("GET")
	router.HandleFunc("/user/{userid}", UserByIdFunc).Methods("GET")
	router.HandleFunc("/game/{gameid}", GameByIdFunc).Methods("GET")
	router.HandleFunc("/video/{videoid}", VideoByIdFunc).Methods("GET")
	
	handler := c.Handler(router)
	handler = Logger(handler, "String")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
