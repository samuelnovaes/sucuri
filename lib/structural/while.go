package structural

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/evaluator"
)

func While(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	condition := args[0].(*ast.Function)
	callback := args[1].(*ast.Function)
	checkCondition := func() bool {
		result, _, _ := evaluator.EvalFunction(*condition, ctx, &ast.Null{})
		return result.(*ast.Boolean).Value
	}
	for checkCondition() {
		result, _, returned := evaluator.EvalFunction(*callback, ctx, &ast.Null{})
		if result.GetKind() == ast.BREAK {
			break
		}
		if result.GetKind() == ast.CONTINUE {
			continue
		}
		if returned {
			return &ast.Return{Value: result}
		}
	}
	return &ast.Null{}
}
