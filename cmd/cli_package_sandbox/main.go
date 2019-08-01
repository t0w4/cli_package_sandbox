package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "cli_package_sandbox"
	app.Usage = "cli package test"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) error {
		fmt.Println("cli test is done!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
