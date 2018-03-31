package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"

	"github.com/gorilla/mux"
)

func ErrorHandle(err, errorOn){
	if err=! nil{
		log.Fatal(errorOn, err)
		return
	}
}

func DataRequest( idName,url ){
	client := &http.Client{}
	vars := mux.Vars(re)
	id := vars[idName]

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Client-ID", "hq22yaqr8c24w8ymv4q0vhh9wrxzst")
	req.Header.Add("Content-Type", "application/json")
	ErrorHandle(err,"NewRequest")

	resp, err := client.Do(req)
	ErrorHandle(err,"Response: ")

	defer resp.Body.Close()
	
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	
    bodyString := string(bodyBytes)
    respondWithJSON(w,http.StatusOK, bodyString)
}

func VideoByIDHander(w http.ResponseWriter, re *http.Request) {
	
	
	url := "https://api.twitch.tv/helix/videos?id=" + id
	idName := "videoid"
	DataRequest(idName, url)
	
	
	

func UserByIdFunc(w http.ResponseWriter, re *http.Request){
	vars := mux.Vars(re)
	id :=vars["userid"]

	
}

func GameByIdFunc(w http.ResponseWriter, re *http.Request){
	vars := mux.Vars(re)
	id :=vars["gameid"]

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
