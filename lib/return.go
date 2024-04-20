package lib

import (
	"github.com/samuelnovaes/sucuri/ast"
)

func Return(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	return &ast.Return{Value: args[0]}
}
