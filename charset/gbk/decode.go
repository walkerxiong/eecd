package gbk

import (
	"errors"
	"github.com/gbabyX/cced/charset/types"
)

var ErrBytes = errors.New("invalid bytes")

func init() {
	ScanCp936Table()
}

func Decode(data []byte) ([]byte, error) {
	var (
		dst  []types.DWord
		size = 0
	)
	for i, l := 0, len(data); i < l; i += size {
		// the upper byte
		h := data[i]
		if h < 0x80 {
			// acii2 code
			size = 1
			dst = append(dst, types.DWord(h))
		} else if h < 0xff {
			size = 2
			var low byte
			if i+1 < l {
				low = data[i+1]
			} else {
				return nil, ErrBytes
			}
			if low > 0x40 && low < 0xFF {
				var code = types.DWord(h)<<8 + types.DWord(low)
				var index = code - 0x8141
				if code == TransformUnicodeTable[index][0] {
					dst = append(dst, TransformUnicodeTable[index][1])
				} else {
					return nil, ErrBytes
				}
			} else {
				return nil, ErrBytes
			}
		}

	}
	return types.DWordsToBytes(dst), nil
}
