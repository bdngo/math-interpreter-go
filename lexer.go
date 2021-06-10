package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

const DIGITS = "0123456789"

func MakeToken(t TokenType) Token {
	return Token{t, math.Inf(1)}
}

func (lex *Lexer) makeNumberToken() (Token, error) {
	var currNumber string
	var numberToken Token
	decPts := 0

	for ; lex.currChar != "" && (lex.currChar == "." || strings.Contains(DIGITS, lex.currChar)); lex.advance() {
		if lex.currChar == "." {
			decPts += 1
			if decPts > 1 {
				break
			}
		}

		currNumber = strings.Join([]string{currNumber, lex.currChar}, "")
	}

	if strings.HasPrefix(currNumber, ".") {
		currNumber = strings.Join([]string{"0", currNumber}, "")
	}
	if strings.HasSuffix(currNumber, ".") {
		currNumber = strings.Join([]string{currNumber, "0"}, "")
	}

	lexedFloat, err := strconv.ParseFloat(currNumber, 64)
	if err != nil {
		return numberToken, errors.New("error lexing number")
	}
	numberToken.tType = NUMBER
	numberToken.value = lexedFloat
	return numberToken, nil
}

type Lexer struct {
	text     string
	currChar string
}

func (lex *Lexer) advance() {
	if len(lex.text) == 0 {
		lex.currChar = ""
	} else {
		lex.currChar = string(lex.text[0])
		lex.text = lex.text[1:]
	}
}

func MakeLexer(txt string) Lexer {
	lex := Lexer{txt, ""}
	lex.advance()
	return lex
}

func (lex *Lexer) Tokenize() ([]Token, error) {
	res := make([]Token, 0)

	for lex.currChar != "" {
		if unicode.IsSpace([]rune(lex.currChar)[0]) {
			lex.advance()
			continue
		} else if lex.currChar == "." || strings.Contains(DIGITS, lex.currChar) {
			numberToken, err := lex.makeNumberToken()
			if err != nil {
				return nil, err
			}
			res = append(res, numberToken)
			continue
		}
		switch lex.currChar {
		case "+":
			res = append(res, MakeToken(PLUS))
		case "-":
			res = append(res, MakeToken(MINUS))
		case "*":
			res = append(res, MakeToken(MULTIPLY))
		case "/":
			res = append(res, MakeToken(DIVIDE))
		case "%":
			res = append(res, MakeToken(MODULO))
		case "^":
			res = append(res, MakeToken(POWER))
		case "(":
			res = append(res, MakeToken(L_PAREN))
		case ")":
			res = append(res, MakeToken(R_PAREN))
		default:
			return nil, fmt.Errorf("illegal character: %s", lex.currChar)
		}
		lex.advance()
	}

	return res, nil
}
