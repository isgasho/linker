package internal

// Config defines structure of a linker configuration file
type Config struct {
	Symlinks Symlinks `hcl:"symlink,block"`
}

// Symlinks is needed in order to allow multiple definitions
// of symlink in a configuation file
type Symlinks []Symlink

// Symlink defines how symlink is represented in Go after
// unmarshaling from HCL/JSON
type Symlink struct {
	Source string `hcl:"source,attr"`
	Target string `hcl:"target,attr"`
}
