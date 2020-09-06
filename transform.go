package cced

import (
	"github.com/gbabyX/cced/charset/gbk"
	"github.com/gbabyX/cced/charset/types"
	"github.com/gbabyX/cced/charset/utf16"
	"github.com/gbabyX/cced/charset/utf8"
)

func GbkToUtf8(data []byte) ([]byte, error) {
	u, err := gbk.Decode(data)
	if err != nil {
		return nil, err
	}
	return utf8.Encode(u)
}

func GbkToUtf16(data []byte) ([]byte, error) {
	u, err := gbk.Decode(data)
	if err != nil {
		return nil, err
	}
	return utf16.Encode(u)
}

func Utf16ToUtf8(data []byte) ([]byte, error) {
	u, err := utf16.DecodeDwords(data)
	if err != nil {
		return nil, err
	}
	return utf8.Encode(types.DWordsToBytes(u))
}

func Utf8ToUtf16(data []byte) ([]byte, error) {
	u, err := utf8.DecodeDwords(data)
	if err != nil {
		return nil, err
	}
	return utf16.Encode(types.DWordsToBytes(u))
}
