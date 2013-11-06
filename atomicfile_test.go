package atomicfile_test

import (
	"os"
	"testing"

	"github.com/ParsePlatform/go.atomicfile"
)

func TestSimple(t *testing.T) {
	names := []string{"TestSimple", "/tmp/foo"}
	for _, name := range names {
		defer os.Remove(name)
		f, err := atomicfile.New(name, os.FileMode(0666))
		if err != nil {
			t.Fatal(err)
		}
		f.Write([]byte("foo"))
		if _, err := os.Stat(name); !os.IsNotExist(err) {
			t.Fatal("did not expect file to exist")
		}
		if err := f.Close(); err != nil {
			t.Fatal(err)
		}
		if _, err := os.Stat(name); err != nil {
			t.Fatalf("expected file to exist: %s", err)
		}
	}
}
