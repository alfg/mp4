package mp4

import (
	"fmt"
	"os"

	"github.com/alfg/mp4/atom"
)

// Open opens a file and returns a &File{}.
func Open(path string) (f *atom.File, err error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	f = &atom.File{
		File: file,
	}

	return f, f.Parse()
}
