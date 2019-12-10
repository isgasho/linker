/*
Linker is a declarative symlink manager in the spirit of terraform.
It is meant as a more flexible replacement of GNU stow.
*/
package main

import (
	"log"

	"github.com/domust/linker"
)

func main() {
	if err := linker.Link("main.hcl"); err != nil {
		log.Fatalln(err)
	}
}
