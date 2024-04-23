package structural

import (
	"github.com/samuelnovaes/sucuri/ast"
)

func Return(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	if len(args) > 0 {
		return &ast.Return{Value: args[0]}
	}
	return &ast.Return{Value: &ast.Null{}}
}
