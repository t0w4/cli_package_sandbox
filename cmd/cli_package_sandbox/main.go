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
	//	cli_package_sandbox [global options] command [command options] [arguments...]
	//
	//	VERSION:
	//	0.0.2
	//
	//	COMMANDS:
	//	help, h  Shows a list of commands or help for one command
	//
	//	GLOBAL OPTIONS:
	//	--name value, -n value  user's name (default: "John Doe") [$USER_NAME]
	//	--age value, -a value   user's age (default: 10) [$USER_AGE]
	//	--help, -h              show help
	//	--version, -v           print the version
	var userName string
	var userAge int64

	app := cli.NewApp()
	app.Name = "cli_package_sandbox"
	app.Usage = "cli package test"
	app.Version = "0.0.2"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "name, n",
			Usage:       "user's name",
			EnvVar:      "USER_NAME",
			Hidden:      false,
			Value:       "John Doe",
			Destination: &userName,
		},
		cli.Int64Flag{
			Name:        "age, a",
			Usage:       "user's age",
			EnvVar:      "USER_AGE",
			Hidden:      false,
			Value:       10,
			Destination: &userAge,
		},
	}
	app.Action = func(c *cli.Context) error {
		fmt.Println(c.Args()) // return args list
		fmt.Println(c.NArg()) // return args count
		if userName == "bob" {
			fmt.Println("You are bob!!!")
		} else {
			fmt.Printf("You are not bob. You are %s\n", userName)
		}
		if c.Int64("age")%2 == 0 {
			fmt.Printf("Your age is even[%d]\n", userAge)
		} else {
			fmt.Printf("Your age is odd[%d]\n", userAge)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
