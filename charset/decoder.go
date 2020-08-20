package charset

type Decoder interface {
	Decode([]byte) ([]rune, error)
}
