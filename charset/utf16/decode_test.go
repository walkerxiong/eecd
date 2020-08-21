package utf16

import "testing"

func TestDecode(t *testing.T) {
	var tv = []byte{0xD8, 0x01, 0xDC, 0x37}
	r, e := Decode(tv)
	if e != nil {
		t.Fatal(e)
	}
	// 10437
	t.Logf("%x", r)
}
