package utf16

import "github.com/gbabyX/cced/charset/types"

func Encode(source []byte) ([]byte, error) {
	var (
		dst []types.DWord
		err error
	)
	data := types.BytesToDWords(source)
	for _, code := range data {
		if code < TwoBytesBoundary {
			dst = append(dst, code)
			continue
		}
		vx := code - TwoBytesBoundary
		vh := vx >> 10
		vl := 0x03ff & vx
		w1 := HighCode | vh
		w2 := LowCode | vl
		dst = append(dst, w1<<16|w2)
	}

	return types.DWordsToBytes(dst), err
}
