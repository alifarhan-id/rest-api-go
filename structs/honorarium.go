package structs

import "github.com/jinzhu/gorm"

type Tabler interface {
	TableName() string
}

type Honorarium struct {
	gorm.Model

	Id         int
	Kode       string
	Uraian     string
	Spek       string
	Satuan     string
	Harga      string
	Rekening   string
	Rekening_2 string
	Rekening_3 string
	Kelompok   string
}
