package lib

import (
	"github.com/samuelnovaes/sucuri/ast"
)

func Inc(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	symbol := args[0].(*ast.Identifier).Symbol
	expr := (*ctx)[symbol]
	if expr.GetKind() != ast.NUMBER {
		panic("Invalid INC operation")
	}
	return Set(ctx, args[0], Add(ctx, expr, &ast.Number{Value: 1}))
}
