package main

func Eval(ast Node) (res float64) {
	ast.visit(NodeVisitor{
		func(nn NumNode) { res = nn.node },
		func(an AddNode) { res = Eval(an.nodeL) + Eval(an.nodeR) },
		func(sn SubNode) { res = Eval(sn.nodeL) - Eval(sn.nodeR) },
		func(mn MulNode) { res = Eval(mn.nodeL) * Eval(mn.nodeR) },
		func(dn DivNode) { res = Eval(dn.nodeL) / Eval(dn.nodeR) },
		func(dn PosNode) { res = Eval(dn.node) },
		func(dn NegNode) { res = -Eval(dn.node) },
	})
	return
}
