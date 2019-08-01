package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	//output:
	//
	//NAME:
	//	cli_package_sandbox - cli package test
	//
	//USAGE:
	//	main [global options] command [command options] [arguments...]
	//
	//	VERSION:
	//	0.0.1
	//
	//	COMMANDS:
	//	help, h  Shows a list of commands or help for one command
	//
	//	GLOBAL OPTIONS:
	//	--help, -h     show help
	//	--version, -v  print the version
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
