package atom

import (
	"encoding/binary"
	"fmt"
	"os"
)

const (
	// BoxHeaderSize Size of box header.
	BoxHeaderSize = int64(8)
)

// File defines a file structure.
type File struct {
	*os.File
	Ftyp *FtypBox
	Moov *MoovBox
	Mdat *MdatBox
	Size int64

	IsFragmented bool
}

// Parse reads an MP4 file for atom boxes.
func (f *File) Parse() error {
	info, err := f.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	// fmt.Printf("Filesize: %v \n", info.Size())
	f.Size = info.Size()

	boxes := readBoxes(f, int64(0), f.Size)
	for _, box := range boxes {
		switch box.Name {
		case "ftyp":
			f.Ftyp = &FtypBox{Box: box}
			f.Ftyp.parse()
		case "wide":
			// fmt.Println("found wide")
		case "mdat":
			f.Mdat = &MdatBox{Box: box}
			// No mdat boxes to parse
		case "moov":
			f.Moov = &MoovBox{Box: box}
			f.Moov.parse()

			f.IsFragmented = f.Moov.IsFragmented
		}
	}
	return nil
}

// ReadBoxAt reads a box from an offset.
func (f *File) ReadBoxAt(offset int64) (boxSize uint32, boxType string) {
	buf := f.ReadBytesAt(BoxHeaderSize, offset)
	boxSize = binary.BigEndian.Uint32(buf[0:4])
	boxType = string(buf[4:8])
	return boxSize, boxType
}

// ReadBytesAt reads a box at n and offset.
func (f *File) ReadBytesAt(n int64, offset int64) (word []byte) {
	buf := make([]byte, n)
	if _, error := f.ReadAt(buf, offset); error != nil {
		fmt.Println(error)
		return
	}
	return buf
}

// Box defines an Atom Box structure.
type Box struct {
	Name        string
	Size, Start int64
	File        *File
}

// ReadBoxData reads the box data from an atom box.
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
