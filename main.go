package main

import (
	"fmt"
	"os"

	"github.com/bgentry/speakeasy"
	"github.com/codegangsta/cli"
	"github.com/foomo/htpasswd"
)

func main() {
	var file string
	app := cli.NewApp()
	app.Name = "htpasswd"
	app.Version = "1.0.0"
	app.Usage = "add and remove credentials to/from an htpasswd file"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "file, f",
			Usage:       "Modify credentials in `FILE`",
			Value:       "filevault.htpasswd",
			Destination: &file,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a user credential",
			Action: func(c *cli.Context) {
				name := c.Args().First()
				if len(name) < 1 {
					fmt.Println("a username must be specified")
					return
				}
				var pw1, pw2 string
				var err error
				pw1, err = speakeasy.Ask("Enter password for " + name + ": ")
				if err != nil {
					fmt.Println("Error gathering password:", err)
					return
				}
				pw2, err = speakeasy.Ask("Confirm password: ")
				if err != nil {
					fmt.Println("Error gathering password:", err)
					return
				}
				if pw1 != pw2 {
					fmt.Println("Passwords did not match")
					return
				}
				err = htpasswd.SetPassword(file, name, pw1, htpasswd.HashSHA)
				if err != nil {
					fmt.Println("Error adding user", name+":", err)
					return
				}
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"r"},
			Usage:   "remove credentials for a user",
			Action: func(c *cli.Context) {
				name := c.Args().First()
				if len(name) < 1 {
					fmt.Println("a username must be specified")
					return
				}
				err := htpasswd.RemoveUser(file, name)
				if err != nil {
					fmt.Println("Error removing user", name+":", err)
					return
				}
			},
		},
	}

	app.Run(os.Args)
}
