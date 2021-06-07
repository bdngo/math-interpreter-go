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

func pop(t *[]Token) Token {
	res := (*t)[0]
	*t = (*t)[1:]
	return res
}

func PostfixEval(tokens []Token) float64 {
	stack := make([]Token, 0)
	for _, t := range tokens {
		var a, b float64
		switch t.tType {
		case PLUS:
			a = pop(&stack).value
			b = pop(&stack).value
			stack = append(stack, Token{NUMBER, a + b})
		case MINUS:
			a = pop(&stack).value
			b = pop(&stack).value
			stack = append(stack, Token{NUMBER, a - b})
		case MULTIPLY:
			a = pop(&stack).value
			b = pop(&stack).value
			stack = append(stack, Token{NUMBER, a * b})
		case DIVIDE:
			a = pop(&stack).value
			b = pop(&stack).value
			stack = append(stack, Token{NUMBER, a / b})
		default:
			stack = append(stack, t)
		}
	}
	return stack[0].value
}
