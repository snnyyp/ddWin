package Pseudo

import (
	. "github.com/snnyyp/ddWin/Define"
	"os"
)

type Stdin struct {
	Stdout
}

func (s *Stdin) New() {
	s.file, _ = os.OpenFile("CONIN$", UniversalOpenMode, DefaultPermission)
}
