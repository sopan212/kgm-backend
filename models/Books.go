package models

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Judul    string `json:"judul"`
	Penulis  string `json:"penulis"`
	Penerbit string `json:"penerbit"`
	Halaman  string `json:"halaman"`
	Ukuran   string `json:"ukuran"`
	Harga    string `json:"harga"`
	Isbn     string `json:"isbn"`
	Image    []byte `json:"image"`
}
