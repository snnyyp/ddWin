package define

import "io"

type IfceIoBase interface {
	io.Reader
	io.Writer
	io.Seeker
	io.Closer
}
