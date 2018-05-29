package main

import (
	"fmt"

	"github.com/alfg/mp4/mp4"
)

func main() {
	file, _ := mp4.Open("videos/ToS-4k-1920.mp4")
	file.Close()

	// fmt.Println(file.Ftyp.Name)
	// fmt.Println(file.Ftyp.MajorBrand)
	// fmt.Println(file.Ftyp.MinorVersion)
	// fmt.Println(file.Ftyp.CompatibleBrands)

	fmt.Println(file.Moov.Name, file.Moov.Size)
	fmt.Println(file.Moov.Mvhd.Name)
	fmt.Println(file.Moov.Mvhd.Version)
	fmt.Println(file.Moov.Mvhd.Volume)
}
