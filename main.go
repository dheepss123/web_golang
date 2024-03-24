package main

import (
	"net/http"
	"web_golang/controller"
	"web_golang/server"
)

func main() {
	// Controller
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/biodata/", controller.Index)
	http.HandleFunc("/biodata/index/", controller.Index)
	http.HandleFunc("/add/", controller.Add)
	http.HandleFunc("/edit/", controller.Edit)
	http.HandleFunc("/delete/", controller.Delete)

	// Runing server
	server.Server()
}
