package conversion

import (
	"github.com/gbabyX/conversion/charset/gbk"
	"github.com/gbabyX/conversion/charset/utf8"
)

func GbkToUtf8(data []byte) ([]byte, error) {
	u, err := gbk.Decode(data)
	if err != nil {
		return nil, err
	}
	return utf8.Encode(u)
}
