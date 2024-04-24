package structural

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/evaluator"
)

func While(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	condition := args[0].(*ast.Function)
	callback := args[1].(*ast.Function)
	checkCondition := func() bool {
		result, _ := evaluator.EvalFunction(*condition, ctx)
		return result.(*ast.Boolean).Value
	}
	for checkCondition() {
		evaluator.EvalFunction(*callback, ctx)
	}
	return &ast.Null{}
}
