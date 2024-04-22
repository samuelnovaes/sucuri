package lib

import (
	"github.com/samuelnovaes/sucuri/ast"
)

func Not(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	if args[0].GetKind() == ast.BOOLEAN {
		return &ast.Boolean{Value: !args[0].(*ast.Boolean).Value}
	}
	panic("Invalid NOT operation")
}
