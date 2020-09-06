package utf8

import "github.com/gbabyX/cced/charset/types"

func Decode(data []byte) ([]byte, error) {
	var dst []byte
	var size = 1
	for i, l := 0, len(data); i < l; i += size {
		item := data[i]
		if item <= 0x80 {
			dst = append(dst, item)
			size = 1
		} else if item >= 0xf0 {
			// 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
			b1 := item & 0x07
			b2 := data[i+1] & 0x3f
			b3 := data[i+2] & 0x3f
			b4 := data[i+3] & 0x3f
			dst = append(dst, b2>>4|b1<<2, b3>>2|b2<<4, b3<<6|b4)
			size = 4
		} else if item >= 0xe0 {
			// 1110xxxx 10xxxxxx 10xxxxxx
			b1 := item & 0x0f
			b2 := data[i+1] & 0x3f
			b3 := data[i+2] & 0x3f
			dst = append(dst, b2>>2|b1<<4, b2<<6|b3)
			size = 3
		} else {
			// 110xxxxx 10xxxxxx
			// two bytes
			b1 := item & 0x0f
			b2 := data[i+1] & 0x3f
			dst = append(dst, b1<<6|b2)
			size = 2
		}
	}
	return dst, nil
}

func DecodeDwords(data []byte) ([]types.DWord, error) {
	var dst []types.DWord
	var size = 1
	for i, l := 0, len(data); i < l; i += size {
		item := data[i]
		if item <= 0x80 {
			dst = append(dst, types.DWord(item))
			size = 1
		} else if item >= 0xf0 {
			// 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
			b1 := item & 0x07
			b2 := data[i+1] & 0x3f
			b3 := data[i+2] & 0x3f
			b4 := data[i+3] & 0x3f
			dst = append(dst, types.DWord(b2>>4|b1<<2)<<16|types.DWord(b3>>2|b2<<4)<<8|types.DWord(b3<<6|b4))
			size = 4
		} else if item >= 0xe0 {
			// 1110xxxx 10xxxxxx 10xxxxxx
			b1 := item & 0x0f
			b2 := data[i+1] & 0x3f
			b3 := data[i+2] & 0x3f
			dst = append(dst, types.DWord(b2>>2|b1<<4)<<8|types.DWord(b2<<6|b3))
			size = 3
		} else {
			// 110xxxxx 10xxxxxx
			// two bytes
			b1 := item & 0x0f
			b2 := data[i+1] & 0x3f
			dst = append(dst, types.DWord(b1<<6|b2))
			size = 2
		}
	}
	return dst, nil
}
