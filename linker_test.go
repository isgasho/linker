package linker

import "testing"

func TestLink(t *testing.T) {
	if err := Link("testdata/main.hcl"); err != nil {
		t.Fatal(err)
	}
}
