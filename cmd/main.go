package main

import "leal-technical-test/internal"

func main() {
	app, err := internal.NewServer()
	if err != nil {
		panic(err)
	}

	app.Run()
}
