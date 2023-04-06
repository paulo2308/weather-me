package main

import "module/cmd/app"

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	app.Routes()
	return nil
}