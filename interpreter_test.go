package main

import (
	"math"
	"testing"
)

func TestInterpNumber(t *testing.T) {
	for i := 0; i < 10; i++ {
		ast := NumNode{float64(i)}
		if Eval(ast) != float64(i) {
			t.Error("incorrect expression")
		}
	}

	ast := NumNode{123.456}
	if Eval(ast) != 123.456 {
		t.Error("incorrect expression")
	}
}

func TestInterpAdd(t *testing.T) {
	ast := AddNode{NumNode{1}, NumNode{2}}
	if Eval(ast) != 3.0 {
		t.Error("incorrect expression")
	}
}

func TestInterpSub(t *testing.T) {
	ast := SubNode{NumNode{1}, NumNode{2}}
	if Eval(ast) != -1 {
		t.Error("incorrect expression")
	}
}

func TestInterpMul(t *testing.T) {
	ast := MulNode{NumNode{1}, NumNode{2}}
	if Eval(ast) != 2 {
		t.Error("incorrect expression")
	}
}

func TestInterpDiv(t *testing.T) {
	ast := DivNode{NumNode{1}, NumNode{2}}
	if Eval(ast) != 0.5 {
		t.Error("incorrect expression")
	}
}

func TestInterpPos(t *testing.T) {
	ast := PosNode{NumNode{2}}
	if Eval(ast) != 2 {
		t.Error("incorrect expression")
	}
}

func TestInterpNeg(t *testing.T) {
	ast := NegNode{NumNode{2}}
	if Eval(ast) != -2 {
		t.Error("incorrect expression")
	}
}

func TestDivZero(t *testing.T) {
	ast := DivNode{NumNode{1}, NumNode{0}}
	if !math.IsInf(Eval(ast), 1) {
		t.Error("incorrect expression")
	}
}

func TestInterpFull(t *testing.T) {
	ast := AddNode{
		NumNode{27},
		MulNode{
			SubNode{
				DivNode{
					NumNode{43},
					NumNode{36}},
				NumNode{48}},
			NumNode{51}}}
	res := Eval(ast)
	if math.Abs(res - -2360.08) > 0.01 {
		t.Error("incorrect expression")
	}
}
