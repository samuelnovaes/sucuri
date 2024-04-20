package lib

import (
	"github.com/samuelnovaes/sucuri/ast"
)

func Set(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	symbol := args[0].(*ast.Identifier).Symbol
	(*ctx)[symbol] = args[1]
	return &ast.Null{}
}
