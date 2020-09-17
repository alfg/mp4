package mp4

import (
	"testing"
)

const testFile = "test/tears-of-steel.mp4"

func TestOpen(t *testing.T) {
	file, err := Open(testFile)
	if err != nil {
		t.Error(err)
	}

	if file.Ftyp.Name != "ftyp" {
		t.Error()
	}

	if file.Ftyp.MajorBrand != "isom" {
		t.Error()
	}
}

func BenchmarkOpen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Open(testFile)
		if err != nil {
			b.Error()
		}
	}
}
