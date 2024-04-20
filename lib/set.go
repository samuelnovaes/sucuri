package lib

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/context"
)

func Set(ctx *context.Context, args ...ast.Expression) ast.Expression {
	symbol := args[0].(*ast.Identifier).Symbol
	(*ctx)[symbol] = args[1]
	return &ast.Null{}
}
