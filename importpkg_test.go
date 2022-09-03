package importpkg_test

import (
	"fmt"
	"testing"

	"github.com/mike-kirpa/importpkg"
)

func TestSummPrint(t *testing.T) {
	want := "Hello, GO"
	get := importpkg.Summ()
	if want != get {
		fmt.Printf("wanted %s, but get %s", want, get)
	}
}
