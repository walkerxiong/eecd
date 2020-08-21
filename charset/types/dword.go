package types

func BytesToDWords(data []byte) []DWord {
	l := len(data)
	if l == 0 {
		return nil
	}
	var dst = make([]DWord, 0, int(l/4)+1)
	align := l % 4
	switch align {
	case 1:
		dst = append(dst, DWord(data[0]))
	case 2:
		dst = append(dst, DWord(data[0])<<8|DWord(data[0]))
	case 3:
		dst = append(dst, DWord(data[0])<<16|DWord(data[1])<<8|DWord(data[2]))
	}
	for i := align; i < len(data); i += 4 {
		v := DWord(data[i])<<24 | DWord(data[i+1])<<16 | DWord(data[i+2])<<8 | DWord(data[i+3])
		dst = append(dst, v)
	}
	return dst
}

func DWordsToBytes(dw []DWord) []byte {
	l := len(dw)
	if l == 0 {
		return nil
	}
	var dst = make([]byte, 0, l*4)
	dst = append(dst, DwordToBytes(dw[0])...)
	for i := 1; i < l; i++ {
		dst = append(dst, byte(dw[i]>>24), byte(dw[i]>>16), byte(dw[i]>>8), byte(dw[i]))
	}
	return dst
}

func DwordToBytes(w DWord) []byte {
	var dst = []byte{byte(w >> 24), byte(w >> 16), byte(w >> 8), byte(w)}
	if w > 0xffffff {
		return dst
	} else if w > 0xffff {
		return dst[0:]
	} else if w > 0xff {
		return dst[1:]
	} else {
		return dst[2:]
	}
}
