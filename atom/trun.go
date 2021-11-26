package atom

import (
	"encoding/binary"
)

// TrunBox - Track Fragment Run Box
// Box Type: trun
// Container: Track Fragment Box (traf)
// Mandatory: No
// Quantity: Zero or more.
type TrunBox struct {
	*Box
	Version     byte
	Flags       uint32
	SampleCount uint32
	//TODO Optional fields not yet implemented
}

func (b *TrunBox) parse() error {
	data := b.ReadBoxData()
	b.Version = data[0]
	// b.Flags = [3]byte{data[1], data[2], data[3]}
	b.Flags = binary.BigEndian.Uint32(data[0:4])
	b.SampleCount = binary.BigEndian.Uint32(data[4:8])
	return nil
}
