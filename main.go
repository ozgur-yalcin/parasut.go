package main

import (
	"parasut/config"
	"parasut/src"
)

func init() {
	config.CompanyID = ""    // Paraşüt tarafından belirlenen firma numarasını yazınız
	config.ClientID = ""     // Paraşüt tarafından belirlenen müşteri numarasını yazınız
	config.ClientSecret = "" // Paraşüt tarafından belirlenen müşteri anahtarını yazınız
	config.Username = ""     // Paraşüte giriş yaparken kullandığınız kullanıcı adını yazınız
	config.Password = ""     // Paraşüte giriş yaparken kullandığınız şifreyi yazınız
}

func main() {
	api := parasut.API{}
	auth := api.Authorize()
	if auth {
	}
}
