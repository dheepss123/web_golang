package controller

import (
	"net/http"
	"text/template"
)

func Index(response http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("view/biodata/biodata.html")

	if err != nil{
		panic(err)
	}

	temp.Execute(response,nil)
}
func Add(response http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("view/biodata/add.html")

	if err != nil{
		panic(err)
	}

	temp.Execute(response,nil)
}
func Edit(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("ini adalah halaman edit"))
}
func Delete(response http.ResponseWriter, request *http.Request) {

}
