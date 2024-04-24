package assignment

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/lib/arithmetic"
	"github.com/samuelnovaes/sucuri/lib/pointer"
)

func Inc(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	return Set(ctx, args[0], arithmetic.Add(ctx, pointer.EvalPtr(ctx, args[0]), &ast.Number{Value: 1}))
}
