package importpkg_test

import (
	"testing"
	"github.com/mike-kirpa/importpkg"
)

func TestSummPrint(t *testing.T) {
	want := "Hello, GO!"
	expected := importpkg.summ()
}
