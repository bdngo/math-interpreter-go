package main

type NodeType int

type Node interface {
	visit(v NodeVisitor)
}

type NodeVisitor struct {
	visitNum func(NumNode)
	visitAdd func(AddNode)
	visitSub func(SubNode)
	visitMul func(MulNode)
	visitDiv func(DivNode)
	visitPos func(PosNode)
	visitNeg func(NegNode)
}

type NumNode struct {
	node float64
}

func (n NumNode) visit(v NodeVisitor) {
	v.visitNum(n)
}

type AddNode struct {
	nodeL Node
	nodeR Node
}

func (a AddNode) visit(v NodeVisitor) {
	v.visitAdd(a)
}

type SubNode struct {
	nodeL Node
	nodeR Node
}

func (s SubNode) visit(v NodeVisitor) {
	v.visitSub(s)
}

type MulNode struct {
	nodeL Node
	nodeR Node
}

func (m MulNode) visit(v NodeVisitor) {
	v.visitMul(m)
}

type DivNode struct {
	nodeL Node
	nodeR Node
}

func (d DivNode) visit(v NodeVisitor) {
	v.visitDiv(d)
}

type PosNode struct {
	node Node
}

func (p PosNode) visit(v NodeVisitor) {
	v.visitPos(p)
}

type NegNode struct {
	node Node
}

func (n NegNode) visit(v NodeVisitor) {
	v.visitNeg(n)
}
