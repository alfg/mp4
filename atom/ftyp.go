package atom

import (
	"encoding/binary"
)

// FtypBox defines the ftyp box.
type FtypBox struct {
	*Box
	MajorBrand       string
	MinorVersion     uint32
	CompatibleBrands []string
}

func (b *FtypBox) parse() error {
	data := b.ReadBoxData()
	b.MajorBrand, b.MinorVersion = string(data[0:4]), binary.BigEndian.Uint32(data[4:8])
	if len(data) > 8 {
		for i := 8; i < len(data); i += 4 {
			b.CompatibleBrands = append(b.CompatibleBrands, string(data[i:i+4]))
		}
	}
	return nil
}
