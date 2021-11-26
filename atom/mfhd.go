package atom

import "encoding/binary"

// MfhdBox - Movie Fragment Header Box
// Box Type: mfhd
// Container: Movie Fragment Box (moof)
// Mandatory: Yes
// Quantity: Exactly one.
//
// The movie fragment header contains a sequence number, as a safety check. The sequence number
// usually starts at 1 and increases for each movie fragment in the file, in the order in which they occur.
// This allows readers to verify integrity of the sequence in environments where undesired re‐ordering
// might occur.
type MfhdBox struct {
	*Box
	Version        byte
	Flags          uint32
	SequenceNumber uint32
}

func (b *MfhdBox) parse() error {
	data := b.ReadBoxData()
	b.Version = data[0]
	b.Flags = binary.BigEndian.Uint32(data[0:4])
	b.SequenceNumber = binary.BigEndian.Uint32(data[4:8])

	return nil
}
