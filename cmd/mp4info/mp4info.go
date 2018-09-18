package main

import (
	"flag"
	"fmt"
	"math"
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
	fmt.Printf("  file Size:	%d\n", f.Size)
	fmt.Printf("  brands:	%s, %s\n\n", f.Ftyp.MajorBrand, f.Ftyp.CompatibleBrands)

	fmt.Println("Movie:")
	fmt.Printf("  duration:	%d ms / %d (%s)\n",
		f.Moov.Mvhd.Duration, f.Moov.Mvhd.Timescale,
		mp4.GetDurationString(f.Moov.Mvhd.Duration, f.Moov.Mvhd.Timescale))
	fmt.Printf("  fragments:	%t\n", f.IsFragmented)
	fmt.Printf("  timescale:	%d\n\n", f.Moov.Mvhd.Timescale)

	fmt.Printf("Found %d Tracks\n\n", len(f.Moov.Traks))

	for _, trak := range f.Moov.Traks {
		fmt.Printf("Track %d:\n", trak.Tkhd.TrackID)
		fmt.Printf("  flags:	%d %s\n", trak.Tkhd.Flags, getFlags(trak.Tkhd.Flags))
		fmt.Printf("  id:		%d\n", trak.Tkhd.TrackID)
		fmt.Printf("  type:		%s\n", getHandlerType(trak.Mdia.Hdlr.Handler))
		fmt.Printf("  duration:	%d ms\n", trak.Tkhd.Duration)
		fmt.Printf("  language:	%s\n", trak.Mdia.Mdhd.LanguageString)
		fmt.Printf("  width:	%d\n", trak.Tkhd.Width/(1<<16))
		fmt.Printf("  height:	%d\n", trak.Tkhd.Height/(1<<16))

		fmt.Println("  media:")
		fmt.Printf("    sample count:	%d\n", trak.Mdia.Minf.Stbl.Stts.SampleCounts[0])
		fmt.Printf("    timescale:		%d\n", trak.Mdia.Mdhd.Timescale)
		fmt.Printf("    duration:		%d (media timescale units)\n", trak.Mdia.Mdhd.Duration)
		fmt.Printf("    duration:		%02.0f (ms)\n\n", math.Floor(float64(trak.Mdia.Mdhd.Duration)/float64(trak.Mdia.Mdhd.Timescale)*1000))
	}

}

func getHandlerType(handler string) string {
	var t string
	if handler == "vide" {
		t = "Video"
	} else if handler == "soun" {
		t = "Sound"
	}
	return t
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
