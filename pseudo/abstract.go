package Pseudo

import "github.com/snnyyp/ddWin/Define"

type pseudoAbstract interface {
	Define.IfceBasicFileHandle
	New()
}
