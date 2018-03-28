package main

import (
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
	handler := c.Handler(router)

	m := &http.ServeMux{}
	//m.HandleFunc("/stream/top", TopStreamHandler)
	m.HandleFunc("/video/{videoid}", VideoByIDHander)
	// m.Handle("/user/{userid}", UserByIdFunc)
	// m.Handle("/game/{gameid}", GameByIdFunc)
	http.Handle("/", m)

	http.ListenAndServe(":8080", handler)
}
