package pseudo

type Zero struct {
	Null
}

func (*Zero) Read(p []byte) (int, error) {
	p = make([]byte, len(p))
	return len(p), nil
}
