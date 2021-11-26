package atom

// MoofBox - Movie Fragment Box
// Box Type: moof
// Container: File
// Mandatory: No
// Quantity: Zero or more.
//
// The movie fragments extend the presentation in time. They provide the information that would
// previously have been in the Movie Box. The actual samples are in Media Data Boxes, as usual, if they are
// in the same file. The data reference index is in the sample description, so it is possible to build
// incremental presentations where the media data is in files other than the file containing the Movie Box.
type MoofBox struct {
	*Box
	Mfhd *MfhdBox
	Traf []*TrafBox
}

func (b *MoofBox) parse() error {
	boxes := readBoxes(b.Reader, b.Start+BoxHeaderSize, b.Size-BoxHeaderSize)

	for _, box := range boxes {
		switch box.Name {
		case "mfhd":
			b.Mfhd = &MfhdBox{Box: box}
			b.Mfhd.parse()

		case "traf":
			traf := &TrafBox{Box: box}
			traf.parse()
			b.Traf = append(b.Traf, traf)
		}
	}
	return nil
}
