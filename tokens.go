package main

type TokenType int

const (
	NUMBER TokenType = iota
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	MODULO
	POWER
	L_PAREN
	R_PAREN
)

type Token struct {
	tType TokenType
	value float64
}
