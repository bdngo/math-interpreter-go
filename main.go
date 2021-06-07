package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	isPostfix := flag.Bool("postfix", false, "whether the calculator is run in postfix or infix mode")
	flag.Parse()

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
		if len(tokens) == 0 {
			continue
		}

		var res float64
		if *isPostfix {
			res = PostfixEval(tokens)
		} else {
			parser := MakeParser(tokens)
			ast, err := parser.parse()
			if err != nil {
				fmt.Println(err)
			}
			res = Eval(ast)
		}

		fmt.Println(res)
	}
}
