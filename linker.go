/*
Package linker allows the users to extend the linker tool
in any way they like without modifying the original one.
This way an original one can be kept small and simple.
*/
package linker

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/domust/linker/internal"
	hcl "github.com/hashicorp/hcl/v2/hclsimple"
)

// Link creates symbolic links based on contents of the file
// passed to it. File must be in either HCL or JSON format.
func Link(filename string) error {
	var config internal.Config
	if err := hcl.DecodeFile(filename, nil, &config); err != nil {
		return err
	}

	base, err := os.Getwd()
	if err != nil {
		return err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	for _, sym := range config.Symlinks {
		src := sym.Source
		tgt := sym.Target
		if !strings.HasPrefix(src, "/") {
			src = path.Join(base, src)
		}

		if strings.HasPrefix(tgt, "~/") {
			tgt = path.Join(home, strings.TrimPrefix(tgt, "~/"))
		} else if !strings.HasPrefix(tgt, "/") {
			tgt = path.Join(base, tgt)
		}

		if linkExists(tgt) {
			continue
		}
		if err := os.Symlink(src, tgt); err != nil {
			return err
		}
		fmt.Printf("Symlink %q successfully created\n", fmt.Sprintf("%s -> %s", tgt, src))
	}
	return nil
}

// linkExists returns true when symlink exists
func linkExists(filename string) bool {
	_, err := os.Lstat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
