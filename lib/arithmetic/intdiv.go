package arithmetic

import (
	"math"

	"github.com/samuelnovaes/sucuri/ast"
)

func IntDiv(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	if args[0].GetKind() == ast.NUMBER && args[1].GetKind() == ast.NUMBER {
		div := args[0].(*ast.Number).Value / args[1].(*ast.Number).Value
		return &ast.Number{Value: math.Floor(div)}
	}
	panic("Invalid DIV operation")
}
