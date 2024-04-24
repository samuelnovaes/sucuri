package structural

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/evaluator"
)

func For(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	initializer := args[0].(*ast.Function)
	condition := args[1].(*ast.Function)
	finalizer := args[2].(*ast.Function)
	callback := args[3].(*ast.Function)
	_, ctxCopy, _ := evaluator.EvalFunction(*initializer, ctx)
	checkCondition := func() bool {
		result, _, _ := evaluator.EvalFunction(*condition, ctxCopy)
		return result.(*ast.Boolean).Value
	}
	for checkCondition() {
		result, _, returned := evaluator.EvalFunction(*callback, ctxCopy)
		if returned {
			return &ast.Return{Value: result}
		}
		evaluator.EvalFunction(*finalizer, ctxCopy)
	}
	return &ast.Null{}
}
