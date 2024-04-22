package lib

import (
	"github.com/samuelnovaes/sucuri/ast"
)

func Sub(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	if args[0].GetKind() != ast.NUMBER {
		panic("Invalid SUB operation")
	}
	if len(args) == 1 {
		return &ast.Number{Value: -args[0].(*ast.Number).Value}
	}
	if args[1].GetKind() != ast.NUMBER {
		panic("Invalid SUB operation")
	}
	return &ast.Number{Value: args[0].(*ast.Number).Value - args[1].(*ast.Number).Value}
}
