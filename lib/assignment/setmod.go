package assignment

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/lib/arithmetic"
)

func SetMod(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	return Set(ctx, args[0], arithmetic.Mod(ctx, args...))
}
