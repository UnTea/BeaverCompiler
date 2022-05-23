package main

import (
	"fmt"
	"github.com/UnTea/Compiler/repl"
	"os"
)

func main() {
	fmt.Printf("This is the Beaver programming language!\n")
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
