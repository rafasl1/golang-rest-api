package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
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

func homePage(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Welcome to the HomePage! :)")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/games", allGames)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}