package structural

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/evaluator"
)

func EvalPtr(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	ref := ast.GetRef(args[0], ctx)
	if ref == nil {
		panic("Inavlid POINTER_VALUE operation")
	}
	return evaluator.EvalExpression(ref.Value, ctx)
}
