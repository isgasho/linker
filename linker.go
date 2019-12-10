/*
Package linker allows the users to extend the linker tool
in any way they like without modifying the original one.
This way an original one can be kept small and simple.
*/
package linker

import (
	"fmt"
	"os"

	"github.com/domust/linker/internal"
	hcl "github.com/hashicorp/hcl/v2/hclsimple"
)

// Link creates symbolic links based on contents of the file
// passed to it. File must be in either HCL or JSON format.
func Link(filename string) error {
	var config internal.Config
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
