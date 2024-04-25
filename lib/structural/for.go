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
	_, ctxCopy, _ := evaluator.EvalFunction(*initializer, ctx, &ast.Null{})
	checkCondition := func() bool {
		result, _, _ := evaluator.EvalFunction(*condition, ctxCopy, &ast.Null{})
		return result.(*ast.Boolean).Value
	}
	finalize := func() {
		evaluator.EvalFunction(*finalizer, ctxCopy, &ast.Null{})
	}
	for checkCondition() {
		result, _, returned := evaluator.EvalFunction(*callback, ctxCopy, &ast.Null{})
		if result.GetKind() == ast.BREAK {
			break
		}
		if result.GetKind() == ast.CONTINUE {
			finalize()
			continue
		}
		if returned {
			return &ast.Return{Value: result}
		}
		finalize()
	}
	return &ast.Null{}
}
