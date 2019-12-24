package atom

import (
	"encoding/binary"
)

type ElstBox struct {
	*Box
	Version    uint32
	EntryCount uint32
	mEntries   []ElstEntry
}

type ElstEntry struct {
	SegmentDuration   uint32
	MediaTime         uint32
	Rate              uint16
	MediaRateFraction uint16
}

func (b *ElstBox) parse() error {
	data := b.ReadBoxData()
	b.Version = binary.BigEndian.Uint32(data[0:4])
	b.EntryCount = binary.BigEndian.Uint32(data[4:8])
	b.mEntries = make([]ElstEntry, b.EntryCount)

	for i := 0; i < len(b.mEntries); i++ {
		b.mEntries[i].SegmentDuration = binary.BigEndian.Uint32(data[(8 + 12*i):(12 + 12*i)])
		b.mEntries[i].MediaTime = binary.BigEndian.Uint32(data[(12 + 12*i):(16 + 12*i)])
		b.mEntries[i].Rate = binary.BigEndian.Uint16(data[(16 + 12*i):(18 + 12*i)])
		b.mEntries[i].MediaRateFraction = binary.BigEndian.Uint16(data[(18 + 12*i):(20 + 12*i)])
	}

	return nil
}

