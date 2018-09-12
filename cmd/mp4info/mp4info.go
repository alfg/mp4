package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alfg/mp4"
)

var input string

func init() {
	flag.StringVar(&input, "i", "", "-i input_file.mp4")
	flag.Parse()
}

func main() {
	if input == "" {
		flag.Usage()
		return
	}

	f, err := mp4.Open(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	defer f.Close()

	// Print out mp4 info.
	fmt.Println("Movie Info")
	fmt.Printf("  File Size: %d\n", f.Size)
	fmt.Printf("  Duration / Timescale: %d / %d (%s)\n",
		f.Moov.Mvhd.Duration, f.Moov.Mvhd.Timescale,
		mp4.GetDurationString(f.Moov.Mvhd.Duration, f.Moov.Mvhd.Timescale))
	fmt.Printf("  Brands: %s, %s\n", f.Ftyp.MajorBrand, f.Ftyp.CompatibleBrands)

	fmt.Println(mp4.GetDurationString(f.Moov.Mvhd.Duration, f.Moov.Mvhd.Timescale))
}
