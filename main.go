package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/beastOP/toy/evaluator"
	"github.com/beastOP/toy/lexer"
	"github.com/beastOP/toy/object"
	"github.com/beastOP/toy/parser"
	"github.com/beastOP/toy/repl"
)

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! We ran into some toy business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func main() {
	args := os.Args
	if len(args) > 1 {
		fileName := args[1]
		data, err := ioutil.ReadFile(fileName)
		env := object.NewEnvironment()
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		dataString := string(data)
		l := lexer.New(dataString)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(os.Stdout, p.Errors())
			return
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(os.Stdout, evaluated.Inspect())
			io.WriteString(os.Stdout, "\n")
		}
	} else {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! This is the toy programming language!\n",
			user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}
}
