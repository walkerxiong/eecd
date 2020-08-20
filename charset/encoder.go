package charset

type Encoder interface {
	Encode([]rune) ([]byte, error)
}
