package lib

import (
	"fmt"

	"github.com/samuelnovaes/sucuri/ast"
)

func String(ctx *ast.Context, args ...ast.Expression) ast.Expression {
	value := fmt.Sprintf("%v", args[0].GetValue(ctx))
	return &ast.String{Value: value}
}
