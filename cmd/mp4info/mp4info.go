package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

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
	fmt.Println("File:")
	fmt.Printf("  file Size: %d\n", f.Size)
	fmt.Printf("  brands: %s, %s\n\n", f.Ftyp.MajorBrand, f.Ftyp.CompatibleBrands)

	fmt.Println("Movie:")
	fmt.Printf("  duration / Timescale: %d / %d (%s)\n",
		f.Moov.Mvhd.Duration, f.Moov.Mvhd.Timescale,
		mp4.GetDurationString(f.Moov.Mvhd.Duration, f.Moov.Mvhd.Timescale))
	fmt.Printf("  fragments: %t\n", f.IsFragmented)
	fmt.Printf("  timescale: %d\n\n", f.Moov.Mvhd.Timescale)

	fmt.Printf("Found %d Tracks\n\n", len(f.Moov.Traks))

	for _, trak := range f.Moov.Traks {
		fmt.Printf("Track %d:\n", trak.Tkhd.TrackID)
		fmt.Printf("  flags: %d %s\n", trak.Tkhd.Flags, getFlags(trak.Tkhd.Flags))
		fmt.Printf("  id: %d\n", trak.Tkhd.TrackID)
		fmt.Printf("  duration: %d\n", trak.Tkhd.Duration)
		fmt.Printf("  width: %d\n", trak.Tkhd.Width/(1<<16))
		fmt.Printf("  height: %d\n", trak.Tkhd.Height/(1<<16))
		fmt.Printf("  type: %s", "test\n") // Check hdlr.
	}
}

func getFlags(flags uint32) string {
	var f []string
	if flags&mp4.TrackFlagEnabled == mp4.TrackFlagEnabled {
		f = append(f, "ENABLED")
	}

	if flags&mp4.TrackFlagInMovie == mp4.TrackFlagInMovie {
		f = append(f, "IN-MOVIE")
	}

	if flags&mp4.TrackFlagInPreview == mp4.TrackFlagInPreview {
		f = append(f, "IN-PREVIEW")
	}
	str := strings.Join(f, " ")
	return str
}
