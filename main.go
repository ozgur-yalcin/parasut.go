package main

import (
	parasut "github.com/OzqurYalcin/parasut/src"
)

func main() {
	config := parasut.Config{CompanyID: "", ClientID: "", ClientSecret: "", Username: "", Password: ""}
	api := &parasut.API{Config: config}
	auth := api.Authorize()
	if auth {

	}
}
