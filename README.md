# golang-rest-api
Simple REST API in Golang Language with REST endpoints that can handle `POST` and `GET` HTTP requests.

![Captura de tela 2021-02-01 191426](https://user-images.githubusercontent.com/49666986/106525139-480a4e00-64c2-11eb-995a-70971713eb7f.png)

Prerequisites: 
- Go version 1.11+ installed on your machine
- <a href="https://github.com/gorilla/mux">Mux</a>, a request router package (you can install this package running `$ go get -u github.com/gorilla/mux` in your terminal)

How to run:<br>
- Open your terminal and enter `$ go run main.go`
- Navigate to `http://localhost:10000/` in your local browser to access the API root
- Navigate to `http://localhost:10000/games` to get the array of games
