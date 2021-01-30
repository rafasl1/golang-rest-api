package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Welcome to the HomePage! :)")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}