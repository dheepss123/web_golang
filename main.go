package main

import (
	"log"
	"net/http"
	"web_golang/config"
	"web_golang/controllers/homecontroler"
)

func main() {
	//connect Database
	config.ConnectDb()

	//1. Homepage
	http.HandleFunc("/", homecontroler.Welcome)

	// 2. Category

	//3. Product
	http.HandleFunc("/product", homecontroler.Welcome)

	//server
	server := http.Server{
		Addr: ":8080",
	}

	log.Println("Server berjalan di 8080")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
