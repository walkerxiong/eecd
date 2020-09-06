package utf16

import (
	"errors"
	"github.com/gbabyX/cced/charset/types"
)

func Decode(data []byte) ([]byte, error) {
	var dst []byte
	dwords := types.BytesToDWords(data)
	if dwords == nil {
		return dst, nil
	}
	for i, l := 0, len(dwords); i < l; i++ {
		v := dwords[i]
		if v < TwoBytesBoundary {
			dst = append(dst, types.DwordToBytes(v)...)
		} else {
			vh := dwords[i] >> 16
			vl := dwords[i] << 16 >> 16
			if vh > HighCode && vl > LowCode {
				w1 := vh - HighCode
				w2 := vl - LowCode
				w := (w1<<10 | w2) + TwoBytesBoundary
				dst = append(dst, types.DwordToBytes(w)...)
			} else {
				return nil, errors.New("err code")
			}
		}
	}
	return dst, nil
}

func DecodeDwords(data []byte) ([]types.DWord, error) {
	var dst []types.DWord
	dwords := types.BytesToDWords(data)
	if dwords == nil {
		return dst, nil
	}
	for i, l := 0, len(dwords); i < l; i++ {
		v := dwords[i]
		if v < TwoBytesBoundary {
			dst = append(dst, v)
		} else {
			vh := dwords[i] >> 16
			vl := dwords[i]
			if vh > HighCode && vl > LowCode {
				w1 := vh - HighCode
				w2 := vl - LowCode
				w := (w1<<10 | w2) + TwoBytesBoundary
				dst = append(dst, w)
			} else {
				return nil, errors.New("err code")
			}
		}
	}
	return dst, nil
}
