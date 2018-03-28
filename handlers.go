package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func VideoByIDHander(w http.ResponseWriter, re *http.Request) {
	client := &http.Client{}
	vars := mux.Vars(re)
	id := vars["videoid"]
	url := "https://api.twitch.tv/helix/videos?id=" + id
	
	fmt.Println("url: ",url)
	req, err := http.NewRequest("GET", url, nil)
	
	req.Header.Add("Client-ID", "hq22yaqr8c24w8ymv4q0vhh9wrxzst")
	fmt.Println("Header: ", req.Header)
	if err != nil {
		fmt.Println("Request error",err)
		log.Fatal("NewRequest: ", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Do error",err)
		log.Fatal("Do: ", err)
		return
	}
	fmt.Println("Response: ",resp)
	
	respondWithJSON(w,http.StatusOK, resp)

	defer resp.Body.Close()

	var record Video
	fmt.Println("Video id: ", record.id)
	fmt.Println("Video Title: ", record.title)
	fmt.Println("user id: ", record.user_id)


	// if err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
}

// func UserByIdFunc(w http.ResponseWriter, re *http.Request){
// 	vars := mux.Vars(re)
// 	id :=vars["userid"]
//
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// }
//
// func GameByIdFunc(w http.ResponseWriter, re *http.Request){
// 	vars := mux.Vars(re)
// 	id :=vars["gameid"]
//
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// }
