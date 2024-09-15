package main

import (
	_ "leal-technical-test/docs"
	"leal-technical-test/internal"
)

// @title mi api
// @version		1.0
// @description	API test.
// @BasePath /
// @securityDefinitions.apikey	ApiKeyAuth
// @in	header
// @name	X-API-Key
// @securityDefinitions.basic	BasicAuth
// @securityDefinitions.apikey	BearerAuth
// @in header
// @name Authorization
func main() {
	app, err := internal.NewServer()
	if err != nil {
		panic(err)
	}

	app.Run()
}
