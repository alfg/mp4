package mp4

import (
	"io"
	"os"

	"github.com/alfg/mp4/atom"
)

// Open opens a file and returns a &File{}.
func Open(path string) (f *atom.File, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	size, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		return nil, err
	} else if _, err := file.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}

	f = &atom.File{
		Closer:        file,
		SectionReader: io.NewSectionReader(file, 0, size),
	}

	return f, f.Parse()
}

func OpenStream(at io.ReaderAt, off, l int64) (f *atom.File, err error) {
	f = &atom.File{
		SectionReader: io.NewSectionReader(at, off, l),
	}

	return f, f.Parse()
}
