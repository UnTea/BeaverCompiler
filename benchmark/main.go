package main

import (
	"flag"
	"fmt"
	"github.com/UnTea/Compiler/compiler"
	"github.com/UnTea/Compiler/evaluator"
	"github.com/UnTea/Compiler/vm"
	"io/ioutil"
	"time"

	"github.com/UnTea/Compiler/lexer"
	"github.com/UnTea/Compiler/object"
	"github.com/UnTea/Compiler/parser"
)

var engine = flag.String("engine", "vm", "use 'vm' or 'eval'")

func main() {
	flag.Parse()

	var duration time.Duration
	var result object.Object

	b, err := ioutil.ReadFile("C:\\Users\\UnT3a\\Documents\\Code\\Go\\src\\github.com\\UnTea\\Compiler\\benchmark\\fibonacci.txt")

	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	fmt.Println(str)

	l := lexer.New(str)

	fmt.Println("================== LEXER START ==================")
	fmt.Println(l)
	fmt.Println("================== LEXER END ==================")

	p := parser.New(l)

	fmt.Println("================== PARSER START ==================")
	fmt.Println(p)
	fmt.Println("================== PARSER END ==================")

	program := p.ParseProgram()

	if *engine == "vm" {
		comp := compiler.New()
		err := comp.Compile(program)

		if err != nil {
			fmt.Printf("compiler error: %s", err)
			return
		}

		machine := vm.New(comp.Bytecode())

		start := time.Now()
		err = machine.Run()

		if err != nil {
			fmt.Printf("vm error: %s", err)
			return
		}

		duration = time.Since(start)
		result = machine.LastPoppedStackElem()
	} else {
		env := object.NewEnvironment()
		start := time.Now()
		result = evaluator.Eval(program, env)
		duration = time.Since(start)
	}

	fmt.Printf(
		"engine=%s, result=%s, duration=%s\n",
		*engine,
		result.Inspect(),
		duration)
}
