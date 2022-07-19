package pseudo

import "crypto/rand"

type Urandom struct {
	Null
}

func (*Urandom) Read(p []byte) (int, error) {
	return rand.Read(p)
}
