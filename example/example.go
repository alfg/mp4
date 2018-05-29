package main

import (
	"fmt"

	"github.com/alfg/mp4/mp4"
)

func main() {
	fmt.Println("hello")

	// file, _ := os.Open("videos/ToS-4k-1920.mp4")
	// // fileInfo, _ := file.Stat()

	// mp4 := mp4.NewDemuxer(file)
	// mp4.Streams()

	// fmt.Println("filesize: ", fileInfo.Size())
	// fmt.Println("duration:")

	file, _ := mp4.Open("videos/ToS-4k-1920.mp4")
	file.Close()
	fmt.Println(file.Ftyp.CompatibleBrands)
}
