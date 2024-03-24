package main

import (
	"net/http"
	"web_golang/controller"
	"web_golang/server"
)

func main() {
	// Controller
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/biodata", controller.Index)
	http.HandleFunc("/biodata/index", controller.Index)
	http.HandleFunc("/biodata/add", controller.Add)
	http.HandleFunc("/biodata/edit", controller.Edit)
	http.HandleFunc("/biodata/delete", controller.Delete)

	// Runing server
	server.Server()
}
