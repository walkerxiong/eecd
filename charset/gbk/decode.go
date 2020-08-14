package gbk

import "errors"

var ErrBytes = errors.New("invalid bytes")

func init() {
	ScanCp936Table()
}

func Decode(data []byte) ([]rune, error) {
	var (
		dst  []rune
		size = 0
	)
	for i, l := 0, len(data); i < l; i += size {
		// the upper byte
		h := data[i]
		if h < 0x80 {
			// acii2 code
			size = 1
			dst = append(dst, rune(h))
		} else if h < 0xff {
			size = 2
			var low byte
			if i+1 < l {
				low = data[i+1]
			} else {
				return nil, ErrBytes
			}
			if low > 0x40 && low < 0xFF {
				var code = uint16(h)<<8 + uint16(low)
				var index = code - 0x8141
				if code == TransformUnicodeTable[index][0] {
					dst = append(dst, rune(TransformUnicodeTable[index][1]))
				} else {
					return nil, ErrBytes
				}
			} else {
				return nil, ErrBytes
			}
		}

	}
	return dst, nil
}
