package types

func BytesToWords(data []byte) []Word {
	l := len(data)
	if l == 0 {
		return nil
	}
	var dst = make([]Word, 0, int(l/2)+1)
	align := l % 2
	switch align {
	case 1:
		dst = append(dst, Word(data[0]))
	}
	for i := align; i < len(data); i += 2 {
		v := Word(data[i]<<8) | Word(data[i+1])
		dst = append(dst, v)
	}
	return dst
}

func WordsToBytes(w []Word) []byte {
	l := len(w)
	if l == 0 {
		return nil
	}
	var dst = make([]byte, 0, l*2)
	dst = append(dst, WordToBytes(w[0])...)
	for i := 1; i < l; i++ {
		dst = append(dst, byte(w[i]>>8), byte(w[i]))
	}
	return dst
}

func WordToBytes(w Word) []byte {
	var dst = []byte{byte(w >> 8), byte(w)}
	if w > 0xff {
		return dst
	} else {
		return dst[0:]
	}
}
