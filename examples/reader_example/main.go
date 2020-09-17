package main

// Example usage of the mp4 package to print out basic mp4 metadata.

import (
	"fmt"
	"os"

	"github.com/alfg/mp4"
)

func main() {

	file, err := os.Open("test/tears-of-steel.mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	size := info.Size()

	mp4, err := mp4.OpenFromReader(file, size)
	if err != nil {
		panic(err)
	}

	fmt.Println(mp4.Ftyp.Name)
	fmt.Println(mp4.Ftyp.MajorBrand)
	fmt.Println(mp4.Ftyp.MinorVersion)
	fmt.Println(mp4.Ftyp.CompatibleBrands)

	fmt.Println(mp4.Moov.Name, mp4.Moov.Size)
	fmt.Println(mp4.Moov.Mvhd.Name)
	fmt.Println(mp4.Moov.Mvhd.Version)
	fmt.Println(mp4.Moov.Mvhd.Volume)

	fmt.Println("trak size: ", mp4.Moov.Traks[0].Size)
	fmt.Println("trak size: ", mp4.Moov.Traks[1].Size)
	fmt.Println("mdat size: ", mp4.Mdat.Size)
}
