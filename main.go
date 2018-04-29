package main

import (
	"parasut/src"
)

func main() {
	api := parasut.API{}
	auth := api.Authorize()

	if auth {
	}

}
