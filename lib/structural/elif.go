package structural

import (
	"github.com/samuelnovaes/sucuri/ast"
)

func Elif(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	if (*ctx)["#IF"] == nil {
		panic("Invalid ELIF operation")
	}
	if (*ctx)["#IF"].Value.(*ast.Boolean).Value {
		return &ast.Null{}
	}
	return If(ctx, args...)
}
