package Pseudo

type Zero struct {
	Null
}

func (*Zero) Read(p []byte) (int, error) {
	p = make([]byte, len(p))
	return len(p), nil
}
