package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Game struct {
	Title string `json:Title`
	Developer string `json:Dev`
	Gender string `json:Gender`
}

type Games []Game

func allGames(response http.ResponseWriter, request *http.Request) {
	games := Games {
		Game { Title: "Bloodborne", Developer: "From Software", Gender: "Action RPG" },
		Game { Title: "Hollow Knight", Developer: "Team Cherry", Gender: "Metroidvania" },
	}

	fmt.Println("Endpoint Hit: All Games Endpoint")
	json.NewEncoder(response).Encode(games)
}

func PostGames(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "POST endpoint worked")
}

func homePage(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Welcome to the HomePage! :)")
	fmt.Println("Endpoint Hit: homePage")
}


func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/games", allGames).Methods("GET")
	myRouter.HandleFunc("/games", PostGames).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}