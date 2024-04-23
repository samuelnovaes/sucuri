package comparison

import "github.com/samuelnovaes/sucuri/ast"

func Lt(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	if args[0].GetKind() == ast.STRING && args[1].GetKind() == ast.STRING {
		return &ast.Boolean{Value: args[0].(*ast.String).Value < args[1].(*ast.String).Value}
	}
	if args[0].GetKind() == ast.NUMBER && args[1].GetKind() == ast.NUMBER {
		return &ast.Boolean{Value: args[0].(*ast.Number).Value < args[1].(*ast.Number).Value}
	}
	panic("Invalid LT operation")
}
