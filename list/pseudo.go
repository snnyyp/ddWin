package list

import (
	"fmt"
)

func ListPseudo() {
	//
	fmt.Println("Pseudo: *urandom*")
	fmt.Println("Implementation: /dev/urandom")
	fmt.Println("Function: pseudorandom number generator")
	fmt.Println()
	//
	fmt.Println("Pseudo: *null*")
	fmt.Println("Implementation: /dev/null")
	fmt.Println("Function: accepts and discards all input; provides an end-of-file indication when read from")
	fmt.Println()
	//
	fmt.Println("Pseudo: *zero*")
	fmt.Println("Implementation: /dev/zero")
	fmt.Println("Function: accepts and discards all input; produces a continuous stream of null characters(zero-value bytes) as output when read from")
	fmt.Println()
	//
	fmt.Println("Pseudo: *stdin*")
	fmt.Println("Implementation: /dev/stdin")
	fmt.Println("Function: standard input stream")
	fmt.Println()
	//
	fmt.Println("Pseudo: *stdout*")
	fmt.Println("Implementation: /dev/stdout")
	fmt.Println("Function: standard output stream")
	fmt.Println()
	//
	fmt.Println("Pseudo: *stderr*")
	fmt.Println("Implementation: /dev/stderr")
	fmt.Println("Function: standard error stream")
	fmt.Println()
}
