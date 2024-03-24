package server

import(
	"net/http"
	"log"
)

func Server() {
	server := http.Server{
		Addr: ":8081",
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	log.Println("Server running in 8081")
}
