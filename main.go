package main

import (
	"server/interfaces/adapter/initialize"
)

// @securityDefinitions.apikey ApiKeyAuth
// @description Enter the token with the `Bearer: ` prefix, e.g. "Bearer abcde12345".
// @in header
// @name Authorization

func main() {
	// Starting the web-framework application
	initialize.App()
}
