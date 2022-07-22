package Define

import "io"

type IfceBasicFileHandle interface {
	io.ReadWriteSeeker
	io.Closer
	Fd() uintptr
}
