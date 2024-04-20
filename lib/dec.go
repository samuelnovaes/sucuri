package lib

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/context"
)

func Dec(ctx *context.Context, args ...ast.Expression) ast.Expression {
	symbol := args[0].(*ast.Identifier).Symbol
	expr := (*ctx)[symbol]
	if expr.GetKind() != ast.NUMBER {
		panic("Invalid DEC operation")
	}
	return Set(ctx, args[0], Sub(ctx, expr, &ast.Number{Value: 1}))
}
