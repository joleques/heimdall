package main

import (
	"os"

	"github.com/heimdall-app/heimdall/src/application"
)

func main() {
	app := application.NewApp(os.Stdout)
	code := app.Run(os.Args[1:])
	os.Exit(code)
}
