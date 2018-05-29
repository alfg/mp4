package mp4

import (
	"encoding/binary"
	"fmt"
	"os"
)

const (
	// BoxHeaderSize Size of box header.
	BoxHeaderSize = int64(8)
)

// Open opens a file and returns a &File{}.
func Open(path string) (f *File, err error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	f = &File{
		File: file,
	}

	return f, f.parse()
}

type File struct {
	*os.File
	Ftyp *FtypBox
	Moov *MoovBox
	size int64
}

func (f *File) parse() error {
	info, err := f.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	// fmt.Printf("Filesize: %v \n", info.Size())
	f.size = info.Size()

	boxes := readBoxes(f, int64(0), f.size)
	for _, box := range boxes {
		switch box.Name {
		case "ftyp":
			f.Ftyp = &FtypBox{Box: box}
			f.Ftyp.parse()
		case "wide":
			// fmt.Println("found wide")
		case "mdat":
			// fmt.Println("found mdat")
		case "moov":
			f.Moov = &MoovBox{Box: box}
			f.Moov.parse()
		}
	}
	return nil
}

func (f *File) ReadBoxAt(offset int64) (boxSize uint32, boxType string) {
	buf := f.ReadBytesAt(BoxHeaderSize, offset)
	boxSize = binary.BigEndian.Uint32(buf[0:4])
	offset += BoxHeaderSize

	boxType = string(buf[4:8])
	return boxSize, boxType
}

func (f *File) ReadBytesAt(n int64, offset int64) (word []byte) {
	buf := make([]byte, n)
	if _, error := f.ReadAt(buf, offset); error != nil {
		fmt.Println(error)
		return
	}
	return buf
}

type Box struct {
	Name        string
	Size, Start int64
	File        *File
}

func (b *Box) ReadBoxData() []byte {
	if b.Size <= BoxHeaderSize {
		return nil
	}
	return b.File.ReadBytesAt(b.Size-BoxHeaderSize, b.Start+BoxHeaderSize)
}

func readBoxes(f *File, start int64, n int64) (l []*Box) {
	for offset := start; offset < start+n; {
		size, name := f.ReadBoxAt(offset)

		b := &Box{
			Name:  string(name),
			Size:  int64(size),
			File:  f,
			Start: offset,
		}

		l = append(l, b)
		offset += int64(size)
	}
	return l
}
