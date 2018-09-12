package main

import (
	"fmt"

	"github.com/alfg/mp4"
)

func main() {
	file, _ := mp4.Open("test/tears-of-steel.mp4")
	file.Close()

	// fmt.Println(file.Ftyp.Name)
	// fmt.Println(file.Ftyp.MajorBrand)
	// fmt.Println(file.Ftyp.MinorVersion)
	// fmt.Println(file.Ftyp.CompatibleBrands)

	fmt.Println(file.Moov.Name, file.Moov.Size)
	fmt.Println(file.Moov.Mvhd.Name)
	fmt.Println(file.Moov.Mvhd.Version)
	fmt.Println(file.Moov.Mvhd.Volume)

	fmt.Println("trak size: ", file.Moov.Traks[0].Size)
	fmt.Println("trak size: ", file.Moov.Traks[1].Size)
	fmt.Println("mdat size: ", file.Mdat.Size)
}
