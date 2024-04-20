package lib

import (
	"github.com/samuelnovaes/sucuri/ast"
)

func Mul(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	if args[0].GetKind() == ast.NUMBER && args[1].GetKind() == ast.NUMBER {
		return &ast.Number{Value: args[0].(*ast.Number).Value * args[1].(*ast.Number).Value}
	}
	panic("Invalid MUL operation")
}
