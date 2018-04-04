package main

import (
"log"
"net/http"
"fmt"

"github.com/gorilla/mux"
"github.com/rs/cors"

)

func main() {
	PORT := ":8080"
	fmt.Println("Initializing Twitch API on port" + PORT)
	
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
	log.Fatal(http.ListenAndServe(PORT, handler))
}
