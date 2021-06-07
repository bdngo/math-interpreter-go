package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("calc > ")
		txt, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		lexer := MakeLexer(txt)
		tokens, err := lexer.Tokenize()
		if err != nil {
			fmt.Println(err)
		}

		parser := MakeParser(tokens)
		ast, err := parser.parse()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(ast)
	}
}
