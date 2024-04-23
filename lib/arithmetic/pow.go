package arithmetic

import (
	"math"

	"github.com/samuelnovaes/sucuri/ast"
)

func Pow(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	if args[0].GetKind() == ast.NUMBER && args[1].GetKind() == ast.NUMBER {
		return &ast.Number{Value: math.Pow(args[0].(*ast.Number).Value, args[1].(*ast.Number).Value)}
	}
	panic("Invalid SUB operation")
}
