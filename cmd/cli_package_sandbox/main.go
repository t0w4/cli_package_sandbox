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
	//	0.0.3
	//
	//	COMMANDS:
	//	status, s  print cli args
	//	user, u    action for user
	//	help, h    Shows a list of commands or help for one command
	//
	//	GLOBAL OPTIONS:
	//	--name NAME, -n NAME  user's NAME (default: "John Doe") [$USER_NAME, $ USERS_NAME]
	//	--age Age, -a Age     user's Age (default: 10) [$USER_AGE, $ USERS_AGE]
	//	--help, -h            show help
	//	--version, -v         print the version
	var userName string
	var userAge int64

	app := cli.NewApp()
	app.Name = "cli_package_sandbox"
	app.Usage = "cli package test"
	app.Version = "0.0.3"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "name, n",
			Usage:       "user's `NAME`",
			EnvVar:      "USER_NAME, USERS_NAME",
			Hidden:      false,
			Value:       "John Doe",
			Destination: &userName,
		},
		cli.Int64Flag{
			Name:        "age, a",
			Usage:       "user's `Age`",
			EnvVar:      "USER_AGE, USERS_AGE",
			Hidden:      false,
			Value:       10,
			Destination: &userAge,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "status",
			Aliases: []string{"s"},
			Usage:   "print cli args",
			Action: func(c *cli.Context) error {
				fmt.Printf("args: %v", os.Args[1:])
				return nil
			},
		},
		{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "action for user",
			Subcommands: []cli.Command{
				{
					Name:  "create",
					Usage: "add a new user",
					Action: func(c *cli.Context) error {
						fmt.Println("user was created")
						return nil
					},
				},
				{
					Name:  "delete",
					Usage: "delete exists user",
					Action: func(c *cli.Context) error {
						fmt.Println("user was deleted")
						return nil
					},
				},
			},
		},
	}
	app.Action = func(c *cli.Context) error {
		fmt.Printf("argss: %v\n", c.Args())      // return args list
		fmt.Printf("arg counts: %v\n", c.NArg()) // return args count
		fmt.Printf("flags: %v\n", app.Flags)
		fmt.Printf("commands: %v\n", app.Commands)
		if userName == "bob" {
			fmt.Println("You are bob!!!")
		} else {
			return cli.NewExitError(fmt.Sprintf("You are not bob. You are %s\n", userName), 1)
		}
		if c.Int64("age")%2 == 0 {
			fmt.Printf("Your age is even[%d]\n", userAge)
		} else {
			return cli.NewExitError(fmt.Sprintf("Your age is odd[%d]\n", userAge), 2)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
