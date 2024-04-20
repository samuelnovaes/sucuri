package lib

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/context"
)

func Mod(ctx *context.Context, args ...ast.Expression) ast.Expression {
	if args[0].GetKind() == ast.NUMBER && args[1].GetKind() == ast.NUMBER {
		a := int(args[0].(*ast.Number).Value)
		b := int(args[1].(*ast.Number).Value)
		return &ast.Number{Value: float64(a % b)}
	}
	panic("Invalid MOD operation")
}
