package structural

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/lib/assignment"
)

func Const(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	symbol := args[0].(*ast.Identifier).Symbol
	(*ctx)[symbol] = &ast.Reference{Value: args[1], Const: false}
	assignment.Set(ctx, args[0], args[1])
	(*ctx)[symbol].Const = true
	return &ast.Null{}
}
