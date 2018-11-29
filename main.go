package main

import (
	"github.com/OzqurYalcin/parasut/config"
	"github.com/OzqurYalcin/parasut/src"
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
	defer api.Unlock()
	auth := api.Authorize()
	if auth {
	}
}
