package Pseudo

import (
	. "github.com/snnyyp/ddWin/Define"
)

func IsPseudo(s string) bool {
	return s == "*stdout*" || s == "*stdin*" || s == "*stderr*" ||
		s == "*null*" || s == "*zero*" || s == "*urandom*"
}

func GetByName(name string) IfceBasicFileHandle {
	var pseudo pseudoAbstract
	switch name {
	case "*stdout*":
		pseudo = &Stdout{}
	case "*stdin*":
		pseudo = &Stdin{}
	case "*stderr*":
		pseudo = &Stderr{}
	case "*null*":
		pseudo = &Null{}
	case "*zero*":
		pseudo = &Zero{}
	case "*urandom*":
		pseudo = &Urandom{}
	}
	pseudo.New()
	return pseudo
}
