package main

import (
	"errors"
	"math"
)

type Parser struct {
	tokens    []Token
	currToken Token
}

func (p *Parser) advance() {
	if len(p.tokens) == 0 {
		p.currToken = MakeToken(NUMBER) // hack to make blank token
	} else {
		p.currToken = p.tokens[0]
		p.tokens = p.tokens[1:]
	}
}

func isParserTokenEmpty(p *Parser) bool {
	return p.currToken.tType == NUMBER && math.IsInf(p.currToken.value, 1)
}

func MakeParser(t []Token) Parser {
	p := Parser{t, MakeToken(NUMBER)}
	p.advance()
	return p
}

func (p *Parser) parse() (Node, error) {
	var dummy Node

	if isParserTokenEmpty(p) {
		return dummy, errors.New("empty token list")
	}

	res, err := p.expr()
	if err != nil {
		return dummy, err
	}

	if len(p.tokens) != 0 {
		return dummy, errors.New("parser not exhausted")
	}

	return res, nil
}

func (p *Parser) expr() (Node, error) {
	res, err := p.term()
	var dummy Node

	if err != nil {
		return dummy, err
	}

	for !isParserTokenEmpty(p) && (p.currToken.tType == PLUS || p.currToken.tType == MINUS) {
		if p.currToken.tType == PLUS {
			p.advance()
			term, err := p.term()
			if err != nil {
				return dummy, err
			}
			res = AddNode{res, term}
		} else if p.currToken.tType == MINUS {
			p.advance()
			term, err := p.term()
			if err != nil {
				return dummy, err
			}
			res = SubNode{res, term}
		}
	}

	return res, nil
}

func (p *Parser) term() (Node, error) {
	res, err := p.exponent()
	var dummy Node

	if err != nil {
		return dummy, err
	}

	for !isParserTokenEmpty(p) && (p.currToken.tType == MULTIPLY || p.currToken.tType == DIVIDE || p.currToken.tType == MODULO) {
		if p.currToken.tType == MULTIPLY {
			p.advance()
			term, err := p.factor()
			if err != nil {
				return dummy, err
			}
			res = MulNode{res, term}
		} else if p.currToken.tType == DIVIDE {
			p.advance()
			term, err := p.factor()
			if err != nil {
				return dummy, err
			}
			res = DivNode{res, term}
		} else if p.currToken.tType == MODULO {
			p.advance()
			term, err := p.factor()
			if err != nil {
				return dummy, err
			}
			res = ModNode{res, term}
		}
	}

	return res, nil
}

func (p *Parser) exponent() (Node, error) {
	res, err := p.factor()
	var dummy Node

	if err != nil {
		return dummy, err
	}

	for !isParserTokenEmpty(p) && p.currToken.tType == POWER {
		if p.currToken.tType == POWER {
			p.advance()
			term, err := p.factor()
			if err != nil {
				return dummy, err
			}
			res = PowNode{res, term}
		}
	}

	return res, nil
}

func (p *Parser) factor() (Node, error) {
	var dummy Node

	token := p.currToken

	switch token.tType {
	case L_PAREN:
		p.advance()
		res, err := p.expr()
		if err != nil {
			return dummy, err
		}

		if p.currToken.tType != R_PAREN {
			return dummy, errors.New("mismatched parentheses")
		}
		p.advance()
		return res, nil
	case NUMBER:
		p.advance()
		return NumNode{token.value}, nil
	case PLUS:
		p.advance()
		tmp, err := p.factor()
		if err != nil {
			return dummy, err
		}
		return PosNode{tmp}, nil
	case MINUS:
		p.advance()
		tmp, err := p.factor()
		if err != nil {
			return dummy, err
		}
		return NegNode{tmp}, nil
	default:
		return dummy, errors.New("bad syntax")
	}
}

func ShuntingYard(tokens []Token) []Token {
	opStack := make([]Token, 0)
	outQueue := make([]Token, 0)

	for _, t := range tokens {
		switch t.tType {
		case NUMBER:
			outQueue = append(outQueue, t)
		case L_PAREN:
			opStack = append(opStack, t)
		case R_PAREN:
			for len(opStack) > 0 && opStack[len(opStack)-1].tType != L_PAREN {
				outQueue = append(outQueue, opStack[len(opStack)-1])
				opStack = opStack[:len(opStack)-1]
			}
			opStack = opStack[:len(opStack)-1]
		default:
			pmSet := map[TokenType]struct{}{
				MULTIPLY: {},
				DIVIDE:   {},
				MODULO:   {},
				POWER:    {},
			}
			if t.tType == PLUS || t.tType == MINUS {
				for len(opStack) > 0 {
					_, ok := pmSet[opStack[len(opStack)-1].tType]
					if ok {
						outQueue = append(outQueue, opStack[len(opStack)-1])
					} else {
						break
					}
				}
				opStack = append(opStack, t)
			} else if t.tType == MULTIPLY || t.tType == DIVIDE {
				for len(opStack) > 0 && opStack[len(opStack)-1].tType == POWER {
					outQueue = append(outQueue, opStack[len(opStack)-1])
				}
				opStack = append(opStack, t)
			}
		}
	}
	// Reversing operator stack
	for i, j := 0, len(opStack)-1; i < j; i, j = i+1, j-1 {
		opStack[i], opStack[j] = opStack[j], opStack[i]
	}
	return append(outQueue, opStack...)
}
