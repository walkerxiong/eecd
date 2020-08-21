package utf16

import "github.com/gbabyX/cced/charset/types"

func Encode(source []byte) ([]byte, error) {
	var (
		dst []types.Word
		err error
	)
	data := types.BytesToDWords(source)
	for _, code := range data {
		if code < TwoBytesBoundary {
			dst = append(dst, types.Word(code))
			continue
		}
		vx := code - TwoBytesBoundary
		vh := types.Word(vx >> 10)
		vl := types.Word(0x03ff & vx)
		w1 := HighCode | vh
		w2 := LowCode | vl
		dst = append(dst, w1)
		dst = append(dst, w2)
	}

	return types.WordsToBytes(dst), err
}
