package lib

import (
	"github.com/samuelnovaes/sucuri/ast"
)

func Or(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	if args[0].GetKind() == ast.BOOLEAN && args[1].GetKind() == ast.BOOLEAN {
		return &ast.Boolean{Value: args[0].(*ast.Boolean).Value || args[1].(*ast.Boolean).Value}
	}
	panic("Invalid DIV operation")
}
