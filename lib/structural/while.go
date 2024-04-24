package structural

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/evaluator"
)

func While(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	condition := args[0].(*ast.Function)
	callback := args[1].(*ast.Function)
	for evaluator.EvalFunction(*condition, ctx).(*ast.Boolean).Value {
		evaluator.EvalFunction(*callback, ctx)
	}
	return &ast.Null{}
}
