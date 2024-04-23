package assignment

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/lib/logical"
)

func SetOr(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	return Set(ctx, args[0], logical.Or(ctx, args...))
}
