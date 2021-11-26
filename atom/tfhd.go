package atom

import (
	"encoding/binary"
)

// TfhdBox - Track Fragment Header Box
// Box Type: tfhd
// Container: Track Fragment Box (traf)
// Mandatory: Yes
// Quantity: Exactly one.
type TfhdBox struct {
	*Box
	Version byte
	Flags   uint32
	TrackID uint32
	//Optional fields
	BaseDataOffset         uint64
	SampleDescriptionIndex uint32
	DefaultSampleDuration  uint32
	DefaultSampleSize      uint32
	DefaultSampleFlags     uint32
}

func (b *TfhdBox) parse() error {
	data := b.ReadBoxData()

	b.Version = data[0]
	b.Flags = binary.BigEndian.Uint32(data[0:4])
	b.TrackID = binary.BigEndian.Uint32(data[4:8])

	oPos := 8
	if b.Flags&1 != 0 {
		b.BaseDataOffset = binary.BigEndian.Uint64(data[16:24])
		oPos += 8
	}
	if b.Flags&2 != 0 {
		b.SampleDescriptionIndex = binary.BigEndian.Uint32(data[oPos : oPos+4])
		oPos += 4
	}
	if b.Flags&8 != 0 {
		b.DefaultSampleDuration = binary.BigEndian.Uint32(data[oPos : oPos+4])
		oPos += 4
	}
	if b.Flags&10 != 0 {
		b.DefaultSampleSize = binary.BigEndian.Uint32(data[oPos : oPos+4])
		oPos += 4
	}
	if b.Flags&20 != 0 {
		b.DefaultSampleFlags = binary.BigEndian.Uint32(data[oPos : oPos+4])
		oPos += 4
	}

	return nil
}
