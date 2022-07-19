package pseudo

import (
	. "github.com/snnyyp/ddWin/define"
	"os"
)

type Stdin struct {
	Stdout
}

func (s *Stdin) New() {
	s.file, _ = os.OpenFile("CONIN$", FileFlag, FilePerm)
}
