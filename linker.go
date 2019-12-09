package linker

import (
	"fmt"
	"os"

	hcl "github.com/hashicorp/hcl/v2/hclsimple"
)

type Config struct {
	Symlinks Symlinks `hcl:"symlink,block"`
}

type Symlinks []Symlink

type Symlink struct {
	Source string `hcl:"source,attr"`
	Target string `hcl:"target,attr"`
}

func Link(filename string) error {
	var config Config
	err := hcl.DecodeFile(filename, nil, &config)
	if err != nil {
		return err
	}

	for _, sym := range config.Symlinks {
		err = os.Symlink(sym.Source, sym.Target)
		if err != nil {
			return err
		}
		fmt.Printf("Symlink %q successfully created\n", fmt.Sprintf("%s -> %s", sym.Target, sym.Source))
	}
	return nil
}
