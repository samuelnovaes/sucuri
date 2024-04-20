package lib

import (
	"github.com/samuelnovaes/sucuri/ast"
)

func Ne(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	isNotEqual := false
	if args[0].GetValue(ctx) != args[1].GetValue(ctx) {
		isNotEqual = true
	}
	return &ast.Boolean{Value: isNotEqual}
}
