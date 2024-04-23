package assignment

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/lib/arithmetic"
)

func Dec(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	return Set(ctx, args[0], arithmetic.Sub(ctx, args[0], &ast.Number{Value: 1}))
}
