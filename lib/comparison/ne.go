package comparison

import (
	"github.com/samuelnovaes/sucuri/ast"
)

func Ne(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	return &ast.Boolean{Value: args[0].GetValue(ctx) != args[1].GetValue(ctx)}
}
