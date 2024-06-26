package structural

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/evaluator"
)

func If(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	condition := args[0].(*ast.Function)
	checkCondition := func() bool {
		result, _, _ := evaluator.EvalFunction(*condition, ctx, &ast.Null{})
		return result.(*ast.Boolean).Value
	}
	var passed bool
	if checkCondition() {
		passed = true
		result, _, returned := evaluator.EvalFunction(*args[1].(*ast.Function), ctx, &ast.Null{})
		if returned {
			return &ast.Return{Value: result}
		}
		if result.GetKind() == ast.BREAK || result.GetKind() == ast.CONTINUE {
			return result
		}
	} else {
		passed = false
	}
	(*ctx)["#IF"] = &ast.Reference{Value: &ast.Boolean{Value: passed}, Const: true}
	return &ast.Null{}
}
