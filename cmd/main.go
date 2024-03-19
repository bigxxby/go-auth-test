package main

import (
	"log"
	"net/http"
	"project/internal/handlers"
)

func main() {
	main, err := handlers.InitMainStruct()
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/auth", main.AuthenticateHandler)
	http.HandleFunc("/refresh", main.RefreshHandler)
	log.Println("Listening and serving on :http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
