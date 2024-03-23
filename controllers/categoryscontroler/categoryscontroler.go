package categoryscontroler

import (
	"html/template"
	"net/http"
	"web_golang/model/categorymodel"
)

func Index(write http.ResponseWriter, request *http.Request) {
	categories := categorymodel.GetAll()
	data := map[string] any{
		"categories": categories,
	}
	temp, err := template.ParseFiles("viwes/category/index.html")
	if err != nil{
		panic(err)
	}
	temp.Execute(write, data)
}

func Add(write http.ResponseWriter, request *http.Request) {

}

func Edit(write http.ResponseWriter, request *http.Request) {

}

func Delete(write htpp.ResponseWriter, request *http.Request) {

}
