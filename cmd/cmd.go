package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/alfg/mp4/mp4"
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
}
