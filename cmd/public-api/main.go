package main

import (
	"os"

	publicapi "github.com/vaskoz/hacktoberfest2023/public-api"
)

// Collect the global variables that we need for the application
var (
	args = os.Args
	out  = os.Stdout
	err  = os.Stderr
	exit = os.Exit
)

func main() {
	httpPort := os.Getenv("PUBLIC_API_HTTP_PORT")
	app := publicapi.NewApp(args, out, err, exit, httpPort)
	app.Run()
}
