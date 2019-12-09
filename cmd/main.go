package main

import "github.com/domust/linker"

func main() {
	if err := linker.Link("main.hcl"); err != nil {
		panic(err)
	}
}
