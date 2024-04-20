package lib

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/context"
)

func Return(ctx *context.Context, args ...ast.Expression) ast.Expression {
	return &ast.Return{Value: args[0]}
}
