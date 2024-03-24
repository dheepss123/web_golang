package controller

import (
	"net/http"
	"text/template"
	"web_golang/entities"
	"web_golang/models"
)

var biodataModel = models.NewBiodataModel()

// Halaman index
func Index(response http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("view/biodata/index.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(response, nil)
}

// Halaman add
func Add(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("view/biodata/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {
		request.ParseForm()

		var biodata entities.Biodata
		biodata.Nama = request.Form.Get("nama_lengkap")
		biodata.Nim = request.Form.Get("nim")
		biodata.Prodi = request.Form.Get("prodi")
		biodata.Divisi = request.Form.Get("divisi")

		// Mengirim data ke database
		biodataModel.Create(biodata)

		// Menyiapkan data untuk dikirim ke template
		data := make(map[string]interface{})
		data["pesan"] = "Data berhasil terkirim"

		temp, _ := template.ParseFiles("view/biodata/add.html")
		temp.Execute(response, data)
	}
}


// Halaman edit
func Edit(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("ini adalah halaman edit"))
}

// Halaman delete
func Delete(response http.ResponseWriter, request *http.Request) {

}
