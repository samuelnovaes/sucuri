package structural

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/evaluator"
)

func Ternary(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	if args[0].GetKind() != ast.BOOLEAN {
		panic("Invalid TERNARY operation")
	}
	var fn ast.Expression
	if args[0].(*ast.Boolean).Value {
		fn = args[1]
	} else {
		fn = args[2]
	}
	result, _, _ := evaluator.EvalFunction(*fn.(*ast.Function), ctx)
	return result
}
