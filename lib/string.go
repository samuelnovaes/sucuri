package lib

import (
	"fmt"

	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/context"
)

func String(ctx *context.Context, args ...ast.Expression) ast.Expression {
	value := fmt.Sprintf("%v", args[0].GetValue())
	return &ast.String{Value: value}
}
