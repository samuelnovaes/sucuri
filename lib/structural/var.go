package structural

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/lib/assignment"
)

func Var(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	if args[0].GetKind() != ast.IDENTIFIER {
		panic("Invalid VAR operation")
	}
	symbol := args[0].(*ast.Identifier).Symbol
	var value ast.Expression = &ast.Null{}
	if len(args) > 1 {
		value = args[1]
	}
	(*ctx)[symbol] = &ast.Reference{Value: value, Const: false}
	return assignment.Set(ctx, args[0], value)
}
