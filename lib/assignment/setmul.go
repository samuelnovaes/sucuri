package assignment

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/lib/arithmetic"
	"github.com/samuelnovaes/sucuri/lib/pointer"
)

func SetMul(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	return Set(ctx, args[0], arithmetic.Mul(ctx, pointer.EvalPtr(ctx, args[0]), args[1]))
}
