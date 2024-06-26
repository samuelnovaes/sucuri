package arithmetic

import (
	"github.com/samuelnovaes/sucuri/ast"
)

func Div(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	if args[0].GetKind() == ast.NUMBER && args[1].GetKind() == ast.NUMBER {
		return &ast.Number{Value: args[0].(*ast.Number).Value / args[1].(*ast.Number).Value}
	}
	panic("Invalid DIV operation")
}
