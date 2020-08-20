package utf16

import (
	"bytes"
	"testing"
)

func TestEncode(t *testing.T) {
	var a = []rune{0x64321}
	b, err := Encode(a)
	if err != nil {
		t.Fatal(err)
	}
	result := []byte{0xD9, 0x50, 0xDF, 0x21}

	t.Log("encode", b)
	if bytes.Compare(b, result) != 0 {
		t.Fatal("encode err")
	}
}
