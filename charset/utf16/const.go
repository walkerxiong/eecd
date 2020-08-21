package utf16

import "github.com/gbabyX/cced/charset/types"

const (
	BigEndian    = 0xFEFF
	LittleEndian = 0xFFFE

	TwoBytesBoundary types.DWord = 0x10000
	HighCode         types.Word  = 0xD800
	LowCode          types.Word  = 0xDC00
)
