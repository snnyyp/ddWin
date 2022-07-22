package DeviceCat

import "fmt"

func PrintPseudoInfo() {
	/*

	 */
	fmt.Println("Available pseudo devices:")
	fmt.Println("  ", "*urandom*:")
	fmt.Println("    ", "Implementation of:", "/dev/urandom")
	fmt.Println("    ", "Function:", "Pseudorandom number generator")
	fmt.Println("  ", "*null*:")
	fmt.Println("    ", "Implementation of:", "/dev/null")
	fmt.Println("    ", "Function:", "Accepts and discards all input; provides an end-of-file "+
		"indication when read from")
	fmt.Println("  ", "*zero*:")
	fmt.Println("    ", "Implementation of:", "/dev/zero")
	fmt.Println("    ", "Function:", "Accepts and discards all input; produces a continuous "+
		"stream of null characters(zero-value bytes) as output when read from")
	fmt.Println("  ", "*stdin*:")
	fmt.Println("    ", "Implementation of:", "/dev/stdin")
	fmt.Println("    ", "Function:", "Standard input stream")
	fmt.Println("  ", "*stdout*:")
	fmt.Println("    ", "Implementation of:", "/dev/stdout")
	fmt.Println("    ", "Function:", "Standard output stream")
	fmt.Println("  ", "*stderr*:")
	fmt.Println("    ", "Implementation of:", "/dev/stderr")
	fmt.Println("    ", "Function:", "Standard error stream")
	fmt.Println()
}
