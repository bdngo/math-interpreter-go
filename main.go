package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	isPostfix := flag.Bool("postfix", false, "evaluate postfix instead of infix expressions")
	printTokens := flag.Bool("t", false, "print tokens from the lexer")
	printAST := flag.Bool("a", false, "print the AST from the parser (infix only)")
	backend := flag.String("backend", "recursive", "{recursive, shunting} backend for parsing infix expressions")
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
		if *printTokens {
			fmt.Println(tokens)
		}

		var res float64
		if *isPostfix {
			res = PostfixEval(tokens)
		} else {
			if *backend == "shunting" {
				pfTokens := ShuntingYard(tokens)
				if *printAST {
					fmt.Println(pfTokens)
				}
				res = PostfixEval(ShuntingYard(tokens))
			} else if *backend == "recursive" {
				parser := MakeParser(tokens)
				ast, err := parser.parse()
				if *printAST {
					fmt.Println(ast)
				}
				if err != nil {
					fmt.Println(err)
				}
				res = Eval(ast)
			}
		}

		fmt.Println(res)
	}
}
