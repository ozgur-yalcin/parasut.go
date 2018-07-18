package main

import (
	"parasut/config"
	"parasut/src"
)

func init() {
	config.CompanyID = ""    // Firma numarası
	config.ClientID = ""     // Müşteri numarası
	config.ClientSecret = "" // Müşteri anahtarı
	config.Username = ""     // Kullanıcı adı
	config.Password = ""     // Şifre
}

func main() {
	api := new(parasut.API)
	api.Lock()
	defer api.Lock()
	auth := api.Authorize()
	if auth {
	}
}
