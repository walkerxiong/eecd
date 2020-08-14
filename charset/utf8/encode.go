package utf8

func Encode(data []rune) ([]byte, error) {
	var dst []byte
	for _, item := range data {
		if item < 0x80 {
			dst = append(dst, byte(item))
		} else if item <= 1<<11-1 {
			// 110xxxxx 10xxxxxx
			b0 := byte(item>>6) | 0xa0
			b1 := byte(item)&0x3f | 0x80
			dst = append(dst, b0, b1)
		} else if item <= 1<<16-1 {
			// 1110xxxx 10xxxxxx 10xxxxxx
			b0 := byte(item>>12) | 0xe0
			b1 := byte(item>>6)&0x3f | 0x80
			b2 := byte(item)&0x3f | 0x80
			dst = append(dst, b0, b1, b2)
		} else {
			// 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
			b0 := byte(item>>18) | 0xf0
			b1 := byte(item>>12)&0x3f | 0x80
			b2 := byte(item>>6)&0x3f | 0x80
			b3 := byte(item)&0x3f | 0x80
			dst = append(dst, b0, b1, b2, b3)
		}
	}
	return dst, nil
}
