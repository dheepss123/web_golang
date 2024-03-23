package homecontroler

import (
	"html/template"
	"net/http"
)

func Welcome(write http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(write, nil)
}
