package atom

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
