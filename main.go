package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

const (
	version = "0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Name = "kpod"
	app.Usage = "Manage pods and images"
	app.Version = version
	app.Author = "Kpod Contributors"
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{}
	app.Commands = []cli.Command{}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
