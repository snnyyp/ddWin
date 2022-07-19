package pseudo

import (
	. "github.com/snnyyp/ddWin/define"
)

func GetPseudoByName(name string) IfceIoBase {
	var pseudo IfcePseudoBase
	switch name {
	case "stdout":
		pseudo = &Stdout{}
	case "stdin":
		pseudo = &Stdin{}
	case "stderr":
		pseudo = &Stderr{}
	case "null":
		pseudo = &Null{}
	case "zero":
		pseudo = &Zero{}
	case "urandom":
		pseudo = &Urandom{}
	}
	pseudo.New()
	return pseudo
}
