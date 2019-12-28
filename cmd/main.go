/*
Linker is a declarative symlink manager in the spirit of terraform.
It is meant as a more flexible replacement of GNU stow.
*/
package main

import (
	"log"
	"os"

	"github.com/domust/linker"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "linker",
		Description: "declarative symbolic manager",
		Commands: []*cli.Command{
			{
				Name:   "link",
				Usage:  "create symbolic links defined in a file",
				Action: link,
			},
			{
				Name:   "unlink",
				Usage:  "remove symbolic links defined in a file",
				Action: unlink,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func link(c *cli.Context) error {
	return linker.Link("main.hcl")
}

func unlink(c *cli.Context) error {
	return linker.Unlink("main.hcl")
}
