// Package atomicfile provides the ability to automically write a file with an
// eventual rename on Close. This allows for a file to always be in a
// consistent state and never represent an in-progress write.
package atomicfile

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// Behaves like os.File, but does an automic rename operation at Close.
type File struct {
	*os.File
	path string
}

// Create a new file that will replace the file at the given path when Closed.
func New(path string, mode os.FileMode) (*File, error) {
	f, err := ioutil.TempFile(filepath.Dir(path), filepath.Base(path))
	if err != nil {
		return nil, err
	}
	os.Chmod(f.Name(), mode)
	return &File{File: f, path: path}, nil
}

// Close the file replacing the configured file.
func (f *File) Close() error {
	if err := f.File.Close(); err != nil {
		return err
	}
	if err := os.Rename(f.Name(), f.path); err != nil {
		return err
	}
	return nil
}
