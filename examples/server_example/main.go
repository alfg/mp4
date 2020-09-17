package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alfg/mp4"
)

type response struct {
	Size             int64    `json:"size"`
	CompatibleBrands []string `json:"compatible_brands"`
	Duration         uint32   `json:"duration"`
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte(`invalid method`))
	}
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(`file required`))
		return
	}
	defer file.Close()

	mp4, err := mp4.OpenFromReader(file, handler.Size)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp := &response{
		Size:             mp4.Size,
		CompatibleBrands: mp4.Ftyp.CompatibleBrands,
		Duration:         mp4.Moov.Mvhd.Duration,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
	return

}

func main() {
	http.HandleFunc("/upload", upload)

	http.ListenAndServe(":8080", nil)
}
