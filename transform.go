package cced

import (
	"github.com/gbabyX/cced/charset/gbk"
	"github.com/gbabyX/cced/charset/utf8"
)

func GbkToUtf8(data []byte) ([]byte, error) {
	u, err := gbk.Decode(data)
	if err != nil {
		return nil, err
	}
	return utf8.Encode(u)
}
