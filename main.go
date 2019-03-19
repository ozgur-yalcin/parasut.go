package main

import (
	parasut "github.com/OzqurYalcin/parasut/src"
)

func main() {
	api := new(parasut.API)
	config := parasut.Config{CompanyID: "", ClientID: "", ClientSecret: "", Username: "", Password: ""}
	auth := api.Authorize(config)
	if auth {

	}
}
