/*
Linker is a declarative symlink manager in the spirit of terraform.
It is meant as a more flexible replacement of GNU stow.
*/
package main

import "github.com/domust/linker"

func main() {
	if err := linker.Link("main.hcl"); err != nil {
		panic(err)
	}
}
