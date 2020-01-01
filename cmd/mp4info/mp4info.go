package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
	"text/template"

	"github.com/alfg/mp4"
	"github.com/alfg/mp4/atom"
)

const tmpl = `
File:
  file size:    {{.Size}}
  brands:       {{.Ftyp.MajorBrand}} {{.Ftyp.CompatibleBrands}}

Movie:
  duration:     {{.Moov.Mvhd.Duration}} ms / {{.Moov.Mvhd.Timescale}} ({{getDurationString .Moov.Mvhd.Duration .Moov.Mvhd.Timescale}})
  fragments:    {{.IsFragmented}}
  timescale:    {{.Moov.Mvhd.Timescale}}

Found {{len .Moov.Traks}} Tracks 
{{range $trak := .Moov.Traks}}
Track: {{$trak.Tkhd.TrackID}}
  flags:    {{$trak.Tkhd.Flags}} {{getFlags $trak.Tkhd.Flags}}
  id:       {{$trak.Tkhd.TrackID}}
  type:     {{getHandlerType $trak.Mdia.Hdlr.Handler}}
  duration: {{$trak.Tkhd.Duration}} ms
  language: {{$trak.Mdia.Mdhd.LanguageString}}
  width:    {{to16 $trak.Tkhd.Width}}
  height:   {{to16 $trak.Tkhd.Height}}
  media:
    sample count:   {{index $trak.Mdia.Minf.Stbl.Stts.SampleCounts 0}}
    timescale:      {{$trak.Mdia.Mdhd.Timescale}}
    duration:       {{$trak.Mdia.Mdhd.Duration}} (media timescale units)
    duration:       {{getDurationMS $trak.Mdia.Mdhd.Duration $trak.Mdia.Mdhd.Timescale}} (ms)
  {{- if (or (ne $trak.Tkhd.GetWidth 0) (ne $trak.Tkhd.GetHeight 0)) }}
    display width:  {{$trak.Tkhd.GetWidth}}
    display width:  {{$trak.Tkhd.GetHeight}}
  {{- end}}
  {{- if (eq (getHandlerType $trak.Mdia.Hdlr.Handler) "Video")}}
    frame rate (computed): {{getFramerate $trak.Mdia.Minf.Stbl.Stts.SampleCounts $trak.Mdia.Mdhd.Duration $trak.Mdia.Mdhd.Timescale}}
  {{- end}}
{{- end}}
`

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

	funcMap := template.FuncMap{
		"getDurationString": atom.GetDurationString,
		"getFlags":          getFlags,
		"getHandlerType":    getHandlerType,
		"to16":              to16,
		"getDurationMS":     getDurationMS,
		"getFramerate":      getFramerate,
	}

	t := template.Must(template.New("tmpl").Funcs(funcMap).Parse(tmpl))
	if err := t.Execute(os.Stdout, f); err != nil {
		panic(err)
	}
}

func getFramerate(sampleCounts []uint32, duration, timescale uint32) string {
	sc := 1000 * sampleCounts[0]
	durationMS := math.Floor(float64(duration) / float64(timescale) * 1000)
	return fmt.Sprintf("%.2f", float64(sc)/durationMS)
}

func getDurationMS(duration, timescale uint32) string {
	return fmt.Sprintf("%.2f", math.Floor(float64(duration)/float64(timescale)*1000))
}

func to16(i atom.Fixed32) int {
	return int(i / (1 << 16))
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
	if flags&atom.TrackFlagEnabled == atom.TrackFlagEnabled {
		f = append(f, "ENABLED")
	}

	if flags&atom.TrackFlagInMovie == atom.TrackFlagInMovie {
		f = append(f, "IN-MOVIE")
	}

	if flags&atom.TrackFlagInPreview == atom.TrackFlagInPreview {
		f = append(f, "IN-PREVIEW")
	}
	str := strings.Join(f, " ")
	return str
}
