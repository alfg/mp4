package atom

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestGetDurationString(t *testing.T) {

	s := GetDurationString(123456789, 12345)

	if s != "02:46:40:549" {
		t.Error()
	}
}

func TestFixed16(t *testing.T) {
	a := []byte{0x00, 0x00}
	b := []byte{0x01, 0x00}

	a1 := fixed16(a)
	b1 := fixed16(b)

	if a1 != 0 {
		t.Error()
	}

	if b1 != 256 {
		fmt.Println(uint16(b1))
		t.Error()
	}

	if uint16(a1) != uint16(binary.BigEndian.Uint16(a)) {
		t.Error()
	}

	if uint16(b1) != uint16(binary.BigEndian.Uint16(b)) {
		t.Error()
	}

}

func TestFixed32(t *testing.T) {
	a := []byte{0x00, 0x01, 0x00, 0x00}

	a1 := fixed32(a)

	if a1 != 65536 {
		t.Error()
	}

	if uint32(a1) != uint32(binary.BigEndian.Uint32(a)) {
		t.Error()
	}
}
