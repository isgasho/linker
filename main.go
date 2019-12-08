package main

import (
	"fmt"
	"os"

	hcl "github.com/hashicorp/hcl/v2/hclsimple"
)

type Config struct {
	Symlink Symlink `hcl:"symlink,block"`
}

type Symlink struct {
	Source string `hcl:"source,attr"`
	Target string `hcl:"target,attr"`
}

func main() {
	var config Config
	err := hcl.DecodeFile("main.hcl", nil, &config)
	if err != nil {
		panic(err)
	}

	err = os.Symlink(config.Symlink.Source, config.Symlink.Target)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Symlink %q successfully created\n", fmt.Sprintf("%s -> %s", config.Symlink.Target, config.Symlink.Source))
}
