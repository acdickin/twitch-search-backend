package main

import (
	"log"
	"net/http"
	"io/ioutil"

	"github.com/gorilla/mux"
)
	
func DataRequest(url){
	
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
	vars := mux.Vars(router)
	id := vars[videoid] 
	url:= "https://api.twitch.tv/helix/videos?id="+id
	DataRequest( url)
}	
	
func UserByIdFunc(w http.ResponseWriter, re *http.Request){
	vars := mux.Vars(router)
	id := vars[userid] 
	url:= "https://api.twitch.tv/helix/user?id="+id
	DataRequest(url)
}

func GameByIdFunc(w http.ResponseWriter, re *http.Request){
	vars := mux.Vars(router)
	id := vars[gameid] 
	url:= "https://api.twitch.tv/helix/games?id="+id
	DataRequest(url)
}

func StreamsByIdFunc(w http.ResponseWriter, re *http.Request){
	url := "https://api.twitch.tv/helix/streams?first=20"
	DataRequest(url)
}

func ErrorHandle(err, errorOn){
	if err != nil{
		log.Fatal(errorOn, err)
		return
	}
}