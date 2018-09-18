package mp4

import (
	"encoding/binary"
	"fmt"
)

// Flag constants.
const (
	TrackFlagEnabled   = 0x0001
	TrackFlagInMovie   = 0x0002
	TrackFlagInPreview = 0x0004
)

// MoovBox defines the moov box structure.
type MoovBox struct {
	*Box
	Mvhd  *MvhdBox
	Traks []*TrakBox

	IsFragmented bool // check for mvex box exists
}

// func readSubBoxes(f *File, start int64, n int64) (boxes chan *Box) {
// 	return readBoxes(f, start+BoxHeaderSize, n-BoxHeaderSize)
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

		case "mvex":
			fmt.Println("found mvex")
			b.IsFragmented = true
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
	Tkhd *TkhdBox
	Mdia *MdiaBox
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
			b.Tkhd = &TkhdBox{Box: box}
			b.Tkhd.parse()

		case "mdia":
			b.Mdia = &MdiaBox{Box: box}
			b.Mdia.parse()

		case "edts":
			// fmt.Println("found edts")
		}
	}
	return nil
}

// TkhdBox defines the track header box structure.
type TkhdBox struct {
	*Box
	Version          byte
	Flags            uint32
	CreationTime     uint32
	ModificationTime uint32
	TrackID          uint32
	Duration         uint32
	Layer            uint16
	AlternateGroup   uint16
	Volume           Fixed16
	Matrix           []byte
	Width, Height    Fixed32
}

func (b *TkhdBox) parse() error {
	data := b.ReadBoxData()
	b.Version = data[0]
	// b.Flags = [3]byte{data[1], data[2], data[3]}
	b.Flags = binary.BigEndian.Uint32(data[0:4])
	b.CreationTime = binary.BigEndian.Uint32(data[4:8])
	b.ModificationTime = binary.BigEndian.Uint32(data[8:12])
	b.TrackID = binary.BigEndian.Uint32(data[12:16])
	b.Duration = binary.BigEndian.Uint32(data[20:24])
	b.Layer = binary.BigEndian.Uint16(data[32:34])
	b.AlternateGroup = binary.BigEndian.Uint16(data[34:36])
	b.Volume = fixed16(data[36:38])
	b.Matrix = data[40:76]
	b.Width = fixed32(data[76:80])
	b.Height = fixed32(data[80:84])
	return nil
}

// MdiaBox defines the mdia box structure.
type MdiaBox struct {
	*Box
	Hdlr *HdlrBox
	Mdhd *MdhdBox
	Minf *MinfBox
}

func (b *MdiaBox) parse() error {
	boxes := readBoxes(b.File, b.Start+BoxHeaderSize, b.Size-BoxHeaderSize)

	for _, box := range boxes {
		switch box.Name {
		case "hdlr":
			b.Hdlr = &HdlrBox{Box: box}
			b.Hdlr.parse()

		case "mdhd":
			b.Mdhd = &MdhdBox{Box: box}
			b.Mdhd.parse()

		case "minf":
			b.Minf = &MinfBox{Box: box}
			b.Minf.parse()
		}
	}
	return nil
}

// HdlrBox defines the hdlr box structure.
type HdlrBox struct {
	*Box
	Version byte
	Flags   uint32
	Handler string
	Name    string
}

func (b *HdlrBox) parse() error {
	data := b.ReadBoxData()
	b.Version = data[0]
	b.Flags = binary.BigEndian.Uint32(data[0:4])
	b.Handler = string(data[8:12])
	b.Name = string(data[24 : b.Size-BoxHeaderSize])
	return nil
}

// MdhdBox defines the mdhd box structure.
type MdhdBox struct {
	*Box
	Version          byte
	Flags            uint32
	CreationTime     uint32
	ModificationTime uint32
	Timescale        uint32
	Duration         uint32
	Language         uint16
	LanguageString   string
}

func (b *MdhdBox) parse() error {
	data := b.ReadBoxData()
	b.Version = data[0]
	b.Flags = binary.BigEndian.Uint32(data[0:4])
	b.CreationTime = binary.BigEndian.Uint32(data[4:8])
	b.ModificationTime = binary.BigEndian.Uint32(data[8:12])
	b.Timescale = binary.BigEndian.Uint32(data[12:16])
	b.Duration = binary.BigEndian.Uint32(data[16:20])
	b.Language = binary.BigEndian.Uint16(data[20:22])
	b.LanguageString = getLanguageString(b.Language)
	return nil
}

func getLanguageString(language uint16) string {
	var lang [3]uint16
	lang[0] = (language >> 10) & 0x1F
	lang[1] = (language >> 5) & 0x1F
	lang[2] = (language) & 0x1F
	return fmt.Sprintf("%s%s%s",
		string(lang[0]+0x60),
		string(lang[1]+0x60),
		string(lang[2]+0x60))
}

// MinfBox defines the minf box structure.
type MinfBox struct {
	*Box
	Vmhd *VmhdBox
	// Dinf *DinfBox
	Stbl *StblBox
}

func (b *MinfBox) parse() error {
	boxes := readBoxes(b.File, b.Start+BoxHeaderSize, b.Size-BoxHeaderSize)

	for _, box := range boxes {
		switch box.Name {
		case "vmhd":
			b.Vmhd = &VmhdBox{Box: box}
			b.Vmhd.parse()

		case "stbl":
			b.Stbl = &StblBox{Box: box}
			b.Stbl.parse()
		}
	}
	return nil
}

// VmhdBox defines the vmhd box structure.
type VmhdBox struct {
	*Box
	Version      byte
	Flags        uint32
	GraphicsMode uint16
	OpColor      uint16
}

func (b *VmhdBox) parse() error {
	data := b.ReadBoxData()
	b.Version = data[0]
	b.Flags = binary.BigEndian.Uint32(data[0:4])
	b.GraphicsMode = binary.BigEndian.Uint16(data[4:6])
	b.OpColor = binary.BigEndian.Uint16(data[6:8])
	return nil
}

// StblBox defines the stbl box structure.
type StblBox struct {
	*Box
	Stts *SttsBox
}

func (b *StblBox) parse() error {
	boxes := readBoxes(b.File, b.Start+BoxHeaderSize, b.Size-BoxHeaderSize)

	for _, box := range boxes {
		switch box.Name {
		case "stts":
			fmt.Println("found stts")
			b.Stts = &SttsBox{Box: box}
			b.Stts.parse()
		}
	}
	return nil
}

// SttsBox defines the stts box structure.
type SttsBox struct {
	*Box
	Version      byte
	Flags        uint32
	EntryCount   uint32
	SampleCounts []uint32
	SampleDeltas []uint32
}

func (b *SttsBox) parse() error {
	data := b.ReadBoxData()

	b.Version = data[0]
	b.Flags = binary.BigEndian.Uint32(data[0:4])

	count := binary.BigEndian.Uint32(data[4:8])
	b.SampleCounts = make([]uint32, count)
	b.SampleDeltas = make([]uint32, count)

	for i := 0; i < int(count); i++ {
		b.SampleCounts[i] = binary.BigEndian.Uint32(data[(8 + 8*i):(12 + 8*i)])
		b.SampleDeltas[i] = binary.BigEndian.Uint32(data[(12 + 8*i):(16 + 8*i)])
	}
	return nil
}
