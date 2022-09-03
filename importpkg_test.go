package importpkg_test

import (
	"testing"

	"github.com/mike-kirpa/importpkg"
)

func TestSummPrint(t *testing.T) {
	want := "Hello, GO"
	get := importpkg.Summ()
	if want != get {
		t.Errorf("wanted %s, but get %s", want, get)
	}
}
