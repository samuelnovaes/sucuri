package assignment

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/lib/logical"
	"github.com/samuelnovaes/sucuri/lib/pointer"
)

func SetAnd(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	return Set(ctx, args[0], logical.And(ctx, pointer.EvalPtr(ctx, args[0]), args[1]))
}
