package atom

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

const (
	// BoxHeaderSize Size of box header.
	BoxHeaderSize = int64(8)
)

// Mp4Reader defines an mp4 reader structure.
type Mp4Reader struct {
	Reader io.ReaderAt
	Ftyp   *FtypBox
	Moov   *MoovBox
	Mdat   *MdatBox
	Size   int64

	IsFragmented bool
}

// Parse reads an MP4 reader for atom boxes.
func (m *Mp4Reader) Parse() error {
	if m.Size == 0 {
		if ofile, ok := m.Reader.(*os.File); ok {
			info, err := ofile.Stat()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return err
			}
			m.Size = info.Size()
		}
	}

	boxes := readBoxes(m, int64(0), m.Size)
	for _, box := range boxes {
		switch box.Name {
		case "ftyp":
			m.Ftyp = &FtypBox{Box: box}
			m.Ftyp.parse()
		case "mdat":
			m.Mdat = &MdatBox{Box: box}
		case "moov":
			m.Moov = &MoovBox{Box: box}
			m.Moov.parse()

			m.IsFragmented = m.Moov.IsFragmented
		}
	}
	return nil
}

// ReadBoxAt reads a box from an offset.
func (m *Mp4Reader) ReadBoxAt(offset int64) (boxSize uint32, boxType string) {
	buf := m.ReadBytesAt(BoxHeaderSize, offset)
	boxSize = binary.BigEndian.Uint32(buf[0:4])
	boxType = string(buf[4:8])
	return boxSize, boxType
}

// ReadBytesAt reads a box at n and offset.
func (m *Mp4Reader) ReadBytesAt(n int64, offset int64) (word []byte) {
	buf := make([]byte, n)
	if _, error := m.Reader.ReadAt(buf, offset); error != nil {
		fmt.Println(error)
		return
	}
	return buf
}

// Box defines an Atom Box structure.
type Box struct {
	Name        string
	Size, Start int64
	Reader      *Mp4Reader
}

// ReadBoxData reads the box data from an atom box.
func (b *Box) ReadBoxData() []byte {
	if b.Size <= BoxHeaderSize {
		return nil
	}
	return b.Reader.ReadBytesAt(b.Size-BoxHeaderSize, b.Start+BoxHeaderSize)
}

func readBoxes(m *Mp4Reader, start int64, n int64) (l []*Box) {
	for offset := start; offset < start+n; {
		size, name := m.ReadBoxAt(offset)

		b := &Box{
			Name:   string(name),
			Size:   int64(size),
			Reader: m,
			Start:  offset,
		}

		l = append(l, b)
		offset += int64(size)
	}
	return l
}
