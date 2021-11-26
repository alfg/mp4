package atom

// TrafBox - Track Fragment Box
// Box Type: traf
// Container: Track Fragment Box (traf)
// Mandatory: Yes
// Quantity: Zero or more.
type TrafBox struct {
	*Box
	Tfhd *TfhdBox
	Trun *TrunBox
}

func (b *TrafBox) parse() error {
	boxes := readBoxes(b.Reader, b.Start+BoxHeaderSize, b.Size-BoxHeaderSize)

	for _, box := range boxes {
		switch box.Name {
		case "tfhd":
			b.Tfhd = &TfhdBox{Box: box}
			b.Tfhd.parse()

		case "trun":
			b.Trun = &TrunBox{Box: box}
			b.Trun.parse()
		}
	}
	return nil
}
