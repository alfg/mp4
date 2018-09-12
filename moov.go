package mp4

import (
	"encoding/binary"
	"fmt"
)

// MoovBox defines the moov box structure.
type MoovBox struct {
	*Box
	Mvhd  *MvhdBox
	Traks []*TrakBox
}

// func readSubBoxes(f *File, start int64, n int64) (boxes chan *Box) {
// 	return readBoxes(f, start+BOX_HEADER_SIZE, n-BOX_HEADER_SIZE)
// }

func (b *MoovBox) parse() error {
	// fmt.Println("read subboxes starting from ", b.Start, "with size: ", b.Size)
	boxes := readBoxes(b.File, b.Start+BoxHeaderSize, b.Size-BoxHeaderSize)

	for _, box := range boxes {
		switch box.Name {
		case "mvhd":
			b.Mvhd = &MvhdBox{Box: box}
			b.Mvhd.parse()

		case "iods":
			// fmt.Println("found iods")

		case "trak":
			trak := &TrakBox{Box: box}
			trak.parse()
			b.Traks = append(b.Traks, trak)

		case "udta":
			// fmt.Println("found udta")
		}

	}
	return nil
}

// MvhdBox defines the mvhd box structure.
type MvhdBox struct {
	*Box
	Flags            uint32
	Version          uint8
	CreationTime     uint32
	ModificationTime uint32
	Timescale        uint32
	Duration         uint32
	Rate             Fixed32
	Volume           Fixed16
}

func (b *MvhdBox) parse() error {
	data := b.ReadBoxData()
	b.Version = data[0]
	b.Timescale = binary.BigEndian.Uint32(data[12:16])
	b.Duration = binary.BigEndian.Uint32(data[16:20])
	b.Rate = fixed32(data[20:24])
	b.Volume = fixed16(data[24:26])
	return nil
}

// TrakBox defines the trak box structure.
type TrakBox struct {
	*Box
	// Tkhd *TkhdBox
	// mdia *MdiaBox
	// edts *EdtsBox
	// chunks []Chunk
	// samples []Sample
}

func (b *TrakBox) parse() error {
	// fmt.Println("read subboxes starting from ", b.Start, "with size: ", b.Size)
	boxes := readBoxes(b.File, b.Start+BoxHeaderSize, b.Size-BoxHeaderSize)

	for _, box := range boxes {
		switch box.Name {
		case "tkhd":
			// fmt.Println("found tkhd")

		case "mdia":
			// fmt.Println("found mdia")

		case "edts":
			// fmt.Println("found edts")

		}
		return nil
	}
	return nil
}

// Fixed16 is an 8.8 Fixed Point Decimal notation
type Fixed16 uint16

func (f Fixed16) String() string {
	return fmt.Sprintf("%v", uint16(f)>>8)
}

func fixed16(bytes []byte) Fixed16 {
	return Fixed16(binary.BigEndian.Uint16(bytes))
}

// Fixed32 is a 16.16 Fixed Point Decimal notation
type Fixed32 uint32

func fixed32(bytes []byte) Fixed32 {
	return Fixed32(binary.BigEndian.Uint32(bytes))
}
