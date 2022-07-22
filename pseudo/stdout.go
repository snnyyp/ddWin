package Pseudo

import (
	. "github.com/snnyyp/ddWin/Define"
	"os"
)

type Stdout struct {
	pseudoAbstract
	file IfceBasicFileHandle
}

func (s *Stdout) New() {
	s.file, _ = os.OpenFile("CONOUT$", UniversalOpenMode, DefaultPermission)
}

func (s *Stdout) Read(p []byte) (int, error) {
	return s.file.Read(p)
}

func (s *Stdout) Write(p []byte) (int, error) {
	return s.file.Write(p)
}

func (s *Stdout) Seek(offset int64, whence int) (int64, error) {
	return s.file.Seek(offset, whence)
}

func (s *Stdout) Close() error {
	return s.file.Close()
}

func (*Stdout) Fd() uintptr {
	return 0
}
