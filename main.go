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

	//m.HandleFunc("/stream/top", TopStreamHandler)
	router.HandleFunc("/video/{videoid}", VideoByIDHander)
	// m.Handle("/user/{userid}", UserByIdFunc)
	// m.Handle("/game/{gameid}", GameByIdFunc)
	handler := c.Handler(router)
	handler = Logger(handler, "String")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
