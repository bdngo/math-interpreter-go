package main

import (
	"testing"
)

func TestParseEmpty(t *testing.T) {
	p := MakeParser(make([]Token, 0))
	_, err := p.parse()
	if err == nil {
		t.Error("did not raise error with empty tokens")
	} else {
		t.Log(err)
	}
}

func TestParseNumber(t *testing.T) {
	for i := 0; i < 10; i++ {
		tokens := []Token{{NUMBER, float64(i)}}
		p := MakeParser(tokens)
		ast, err := p.parse()
		if err != nil {
			t.Error("failed to parse", err)
			return
		}
		expected := NumNode{float64(i)}
		if ast != expected {
			t.Error("incorrect AST")
		}
	}

	tokens := []Token{{NUMBER, 123.456}}
	p := MakeParser(tokens)
	ast, err := p.parse()
	if err != nil {
		t.Error("failed to parse", err)
		return
	}
	expected := NumNode{123.456}
	if ast != expected {
		t.Error("incorrect AST")
	}
}

func TestParseAdd(t *testing.T) {
	tokens := []Token{
		{NUMBER, 1},
		MakeToken(PLUS),
		{NUMBER, 2}}
	p := MakeParser(tokens)
	ast, err := p.parse()
	if err != nil {
		t.Error("failed to parse", err)
		return
	}
	expected := AddNode{NumNode{1}, NumNode{2}}
	if ast != expected {
		t.Error("incorrect AST")
	}
}

func TestParseSub(t *testing.T) {
	tokens := []Token{
		{NUMBER, 1},
		MakeToken(MINUS),
		{NUMBER, 2}}
	p := MakeParser(tokens)
	ast, err := p.parse()
	if err != nil {
		t.Error("failed to parse", err)
		return
	}
	expected := SubNode{NumNode{1}, NumNode{2}}
	if ast != expected {
		t.Error("incorrect AST")
	}
}

func TestParseMul(t *testing.T) {
	tokens := []Token{
		{NUMBER, 1},
		MakeToken(MULTIPLY),
		{NUMBER, 2}}
	p := MakeParser(tokens)
	ast, err := p.parse()
	if err != nil {
		t.Error("failed to parse", err)
		return
	}
	expected := MulNode{NumNode{1}, NumNode{2}}
	if ast != expected {
		t.Error("incorrect AST")
	}
}

func TestParseDiv(t *testing.T) {
	tokens := []Token{
		{NUMBER, 1},
		MakeToken(DIVIDE),
		{NUMBER, 2}}
	p := MakeParser(tokens)
	ast, err := p.parse()
	if err != nil {
		t.Error("failed to parse", err)
		return
	}
	expected := DivNode{NumNode{1}, NumNode{2}}
	if ast != expected {
		t.Error("incorrect AST")
	}
}

func TestParsePos(t *testing.T) {
	tokens := []Token{
		MakeToken(PLUS),
		{NUMBER, 2}}
	p := MakeParser(tokens)
	ast, err := p.parse()
	if err != nil {
		t.Error("failed to parse", err)
		return
	}
	expected := PosNode{NumNode{2}}
	if ast != expected {
		t.Error("incorrect AST")
	}
}

func TestParseNeg(t *testing.T) {
	tokens := []Token{
		MakeToken(MINUS),
		{NUMBER, 2}}
	p := MakeParser(tokens)
	ast, err := p.parse()
	if err != nil {
		t.Error("failed to parse", err)
		return
	}
	expected := NegNode{NumNode{2}}
	if ast != expected {
		t.Error("incorrect AST")
	}
}

func TestParseFull(t *testing.T) {
	tokens := []Token{
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
	p := MakeParser(tokens)
	ast, err := p.parse()
	if err != nil {
		t.Error("failed to parse", err)
		return
	}
	expected := AddNode{
		NumNode{27},
		MulNode{
			SubNode{
				DivNode{
					NumNode{43},
					NumNode{36}},
				NumNode{48}},
			NumNode{51}}}
	if ast != expected {
		t.Error("incorrect AST")
	}
}
