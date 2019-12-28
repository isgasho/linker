package linker

import "testing"

func TestLink(t *testing.T) {
	testFile := "testdata/main.hcl"
	defer Unlink(testFile)

	if err := Link(testFile); err != nil {
		t.Fatal(err)
	}
}
