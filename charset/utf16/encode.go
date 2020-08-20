package utf16

func Encode(data []rune) ([]byte, error) {
	var (
		dst []byte
		err error
	)
	for _, code := range data {
		if code <= 0xFFFF {
			dst = append(dst, byte(code>>8), byte(code))
		}
		vx := code - 0x10000
		vh := uint16(vx >> 10)
		vl := uint16(0x03ff & vx)
		w1 := 0xD800 | vh
		w2 := 0xDC00 | vl
		dst = append(dst, byte(w1>>8), byte(w1), byte(w2>>8), byte(w2))
	}

	return dst, err
}
