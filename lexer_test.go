package main

import (
	"fmt"
	"testing"
)

func TestLexEmpty(t *testing.T) {
	lex := MakeLexer("")
	tokens, err := lex.Tokenize()
	if err != nil {
		t.Error("failed to tokenize", err)
		return
	}
	if len(tokens) != 0 {
		t.Error("returned nonzero tokens")
	}
}

func TestLexNumber(t *testing.T) {
	for i := 0; i < 10; i++ {
		lex := MakeLexer(fmt.Sprint(i))
		tokens, err := lex.Tokenize()
		if err != nil {
			t.Error("failed to tokenize", err)
			return
		}
		expected := Token{NUMBER, float64(i)}
		if tokens[0] != expected {
			t.Error("number tokens not equal")
		}
	}

	lex := MakeLexer("123.456")
	tokens, err := lex.Tokenize()
	if err != nil {
		t.Error("failed to tokenize", err)
		return
	}
	expected := Token{NUMBER, 123.456}
	if tokens[0] != expected {
		t.Error("number tokens not equal")
	}
}

func TestLexOperators(t *testing.T) {
	lex := MakeLexer("+-*/")
	tokens, err := lex.Tokenize()
	if err != nil {
		t.Error("failed to tokenize", err)
		return
	}
	expected := []Token{
		MakeToken(PLUS),
		MakeToken(MINUS),
		MakeToken(MULTIPLY),
		MakeToken(DIVIDE)}
	for i := 0; i < len(tokens); i++ {
		if tokens[i] != expected[i] {
			t.Error("incorrect operator token")
		}
	}
}

func TestLexParens(t *testing.T) {
	lex := MakeLexer("()")
	tokens, err := lex.Tokenize()
	if err != nil {
		t.Error("failed to tokenize", err)
		return
	}
	expected := []Token{
		MakeToken(L_PAREN),
		MakeToken(R_PAREN)}
	for i := 0; i < len(tokens); i++ {
		if tokens[i] != expected[i] {
			t.Error("incorrect operator token")
		}
	}
}

func TestLexFull(t *testing.T) {
	lex := MakeLexer("27 + (43 / 36 - 48) * 51")
	tokens, err := lex.Tokenize()
	if err != nil {
		t.Error("failed to tokenize", err)
		return
	}
	expected := []Token{
		{NUMBER, 27.0},
		MakeToken(PLUS),
		MakeToken(L_PAREN),
		{NUMBER, 43.0},
		MakeToken(DIVIDE),
		{NUMBER, 36.0},
		MakeToken(MINUS),
		{NUMBER, 48.0},
		MakeToken(R_PAREN),
		MakeToken(MULTIPLY),
		{NUMBER, 51.0},
	}
	for i := 0; i < len(tokens); i++ {
		if tokens[i] != expected[i] {
			t.Error("incorrect operator token")
		}
	}
}
