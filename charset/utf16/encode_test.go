package utf16

import (
	"bytes"
	"testing"
)

func TestEncode(t *testing.T) {
	var a = []byte{0x06, 0x43, 0x21}
	b, err := Encode(a)
	if err != nil {
		t.Fatal(err)
	}
	result := []byte{0xD9, 0x50, 0xDF, 0x21}

	t.Logf("encode,%x \n", b)
	if bytes.Compare(b, result) != 0 {
		t.Fatal("encode err")
	}
}
