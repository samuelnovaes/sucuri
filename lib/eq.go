package lib

import (
	"github.com/samuelnovaes/sucuri/ast"
)

func Eq(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	isEqual := false
	if args[0].GetValue(ctx) == args[1].GetValue(ctx) {
		isEqual = true
	}
	return &ast.Boolean{Value: isEqual}
}
