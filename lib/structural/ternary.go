package structural

import (
	"github.com/samuelnovaes/sucuri/ast"
)

func Ternary(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	if args[0].GetKind() != ast.BOOLEAN {
		panic("Invalid TERNARY operation")
	}
	if args[0].(*ast.Boolean).Value {
		return args[1]
	}
	return args[2]
}
