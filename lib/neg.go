package lib

import (
	"github.com/samuelnovaes/sucuri/ast"
)

func Neg(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	if args[0].GetKind() == ast.NUMBER {
		return &ast.Number{Value: -args[0].(*ast.Number).Value}
	}
	if args[0].GetKind() == ast.BOOLEAN {
		return &ast.Boolean{Value: !args[0].(*ast.Boolean).Value}
	}
	panic("Invalid DIV operation")
}
