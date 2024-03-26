package entities

type Biodata struct {
	Id     int64  `validate:"required" `
	Nama   string `validate:"required" label:"Nama"` 
	Nim    string `validate:"required" label:"Nim"`
	Prodi  string `validate:"required" label:"Prodi"`
	Divisi string `validate:"required" label:"Divisi"`
}
