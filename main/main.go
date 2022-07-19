package main

import "github.com/snnyyp/ddWin/pseudo"

func main() {
	//device := pseudo.Stdout{}
	//device.New()
	device := pseudo.GetPseudoByName("stdout")
	defer device.Close()
	device.Write([]byte("Hello, world!\n"))
}
