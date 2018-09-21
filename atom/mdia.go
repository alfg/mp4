package atom

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
