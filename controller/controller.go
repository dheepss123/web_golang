package controller

import (
	"net/http"
	"strconv"
	"text/template"
	"web_golang/entities"
	"web_golang/libraries"
	"web_golang/models"
)

var validation = libraries.NewValidation()
var biodataModel = models.NewBiodataModel()

// Halaman index
func Index(response http.ResponseWriter, request *http.Request) {
	biodata, _ := biodataModel.FindAll()
	data := map[string]interface{}{
		"biodata": biodata,
	}
	temp, err := template.ParseFiles("view/biodata/index.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(response, data)
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
		biodata.Nama = request.Form.Get("nama")
		biodata.Nim = request.Form.Get("nim")
		biodata.Prodi = request.Form.Get("prodi")
		biodata.Divisi = request.Form.Get("divisi")
		//validasi data
		var data = make(map[string]interface{})
		vErors := validation.Struct(biodata)

		if vErors != nil {
			data["biodata"] = biodata
			data["validation"] = vErors

		} else {
			data["pesan"] = "Biodata berhasil disimpan"
			biodataModel.Create(biodata)
		}

		// Mengirim data ke database
		biodataModel.Create(biodata)

		// Menyiapkan data untuk dikirim ke template
		temp, _ := template.ParseFiles("view/biodata/add.html")
		temp.Execute(response, data)
	}
}

// Halaman edit
func Edit(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {

		querrystring := request.URL.Query()
		id, _ := strconv.ParseInt(querrystring.Get("id"), 10, 64)
		var biodata entities.Biodata
		biodataModel.Find(id, &biodata)

		data := map[string]interface{}{
			"biodata": biodata,
		}

		temp, err := template.ParseFiles("view/biodata/edit.html")

		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)
	} else if request.Method == http.MethodPost {
		request.ParseForm()

		var biodata entities.Biodata
		biodata.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		biodata.Nama = request.Form.Get("nama")
		biodata.Nim = request.Form.Get("nim")
		biodata.Prodi = request.Form.Get("prodi")
		biodata.Divisi = request.Form.Get("divisi")
		//validasi data
		var data = make(map[string]interface{})
		vErors := validation.Struct(biodata)

		if vErors != nil {
			data["biodata"] = biodata
			data["validation"] = vErors

		} else {
			data["pesan"] = "Biodata berhasil diperbaharui"
			biodataModel.Create(biodata)
		}

		// Mengirim data ke database
		biodataModel.Update(biodata)

		// Menyiapkan data untuk dikirim ke template
		temp, _ := template.ParseFiles("view/biodata/edit.html")
		temp.Execute(response, data)
	}
}

// Halaman delete
func Delete(response http.ResponseWriter, request *http.Request) {
	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	biodataModel.Delete(id)

	http.Redirect(response, request, "/pasien", http.StatusSeeOther)
}
