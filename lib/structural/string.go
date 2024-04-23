package structural

import (
	"github.com/samuelnovaes/sucuri/ast"
)

func String(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	return &ast.String{Value: args[0].ToString(ctx)}
}
