package lib

import (
	"fmt"

	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/context"
)

func Format(ctx *context.Context, args ...ast.Expression) ast.Expression {
	format := args[0].(*ast.String).Value
	strs := []any{}
	for _, arg := range args[1:] {
		strs = append(strs, arg.GetValue())
	}
	return &ast.String{Value: fmt.Sprintf(format, strs...)}
}
