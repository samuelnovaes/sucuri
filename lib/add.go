package lib

import (
	"fmt"

	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/context"
)

func Add(ctx *context.Context, args ...ast.Expression) ast.Expression {
	if args[0].GetKind() == ast.NUMBER && args[1].GetKind() == ast.NUMBER {
		return &ast.Number{Value: args[0].(*ast.Number).Value + args[1].(*ast.Number).Value}
	}
	if args[0].GetKind() == ast.STRING || args[1].GetKind() == ast.STRING {
		return &ast.String{Value: fmt.Sprintf("%v%v", args[0].GetValue(), args[1].GetValue())}
	}
	panic("Invalid ADD operation")
}
