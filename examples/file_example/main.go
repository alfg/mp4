package main

// Example usage of the mp4 package to print out basic mp4 metadata.

import (
	"fmt"

	"github.com/alfg/mp4"
)

func main() {
	mp4, _ := mp4.Open("test/tears-of-steel.mp4")

	fmt.Println(mp4.Ftyp.Name)
	fmt.Println(mp4.Ftyp.MajorBrand)
	fmt.Println(mp4.Ftyp.MinorVersion)
	fmt.Println(mp4.Ftyp.CompatibleBrands)

	fmt.Println(mp4.Moov.Name, mp4.Moov.Size)
	fmt.Println(mp4.Moov.Mvhd.Name)
	fmt.Println(mp4.Moov.Mvhd.Version)
	fmt.Println(mp4.Moov.Mvhd.Volume)

	fmt.Println("trak size: ", mp4.Moov.Traks[0].Size)
	fmt.Println("mdhd language: ", mp4.Moov.Traks[0].Mdia.Mdhd.LanguageString)
	fmt.Println("trak size: ", mp4.Moov.Traks[1].Size)
	fmt.Println("mdat size: ", mp4.Mdat.Size)
}
